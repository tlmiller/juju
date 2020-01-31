// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider

import (
	"fmt"
	"time"

	jujuclock "github.com/juju/clock"
	"github.com/juju/errors"
	"github.com/juju/juju/core/watcher"
	"gopkg.in/juju/worker.v1/catacomb"

	core "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

// kubernetesNotifyWatcher reports changes to kubernetes
// resources. A native kubernetes watcher is passed
// in to generate change events from the kubernetes
// model. These events are consolidated into a Juju
// notification watcher event.
type kubernetesNotifyWatcher struct {
	clock    jujuclock.Clock
	catacomb catacomb.Catacomb

	informer cache.SharedIndexInformer
	name     string
	out      chan struct{}
}

func newKubernetesNotifyWatcher(informer cache.SharedIndexInformer, name string, clock jujuclock.Clock) (*kubernetesNotifyWatcher, error) {
	w := &kubernetesNotifyWatcher{
		clock:    clock,
		informer: informer,
		name:     name,
		out:      make(chan struct{}),
	}
	err := catacomb.Invoke(catacomb.Plan{
		Site: &w.catacomb,
		Work: w.loop,
	})
	return w, err
}

const sendDelay = 1 * time.Second

func (w *kubernetesNotifyWatcher) loop() error {
	signals := make(chan struct{}, 1)
	defer close(w.out)

	w.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			logger.Tracef("received k8s event: Added")
			if pod, ok := obj.(*core.Pod); ok {
				logger.Tracef("%v(%v) = %v, status=%+v", pod.Name, pod.UID, pod.Labels, pod.Status)
			}
			if ns, ok := obj.(*core.Namespace); ok {
				logger.Tracef("%v(%v) = %v, status=%+v", ns.Name, ns.UID, ns.Labels, ns.Status)
			}

			select {
			case signals <- struct{}{}:
			default:
			}
			logger.Debugf("fire notify watcher for %v", w.name)
		},
		DeleteFunc: func(obj interface{}) {
			logger.Tracef("received k8s event: Deleted")
			if pod, ok := obj.(*core.Pod); ok {
				logger.Tracef("%v(%v) = %v, status=%+v", pod.Name, pod.UID, pod.Labels, pod.Status)
			}
			if ns, ok := obj.(*core.Namespace); ok {
				logger.Tracef("%v(%v) = %v, status=%+v", ns.Name, ns.UID, ns.Labels, ns.Status)
			}

			select {
			case signals <- struct{}{}:
			default:
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			logger.Tracef("received k8s event: Updated")

			if pod, ok := newObj.(*core.Pod); ok {
				logger.Tracef("%v(%v) = %v, status=%+v", pod.Name, pod.UID, pod.Labels, pod.Status)
			}
			if ns, ok := newObj.(*core.Namespace); ok {
				logger.Tracef("%v(%v) = %v, status=%+v", ns.Name, ns.UID, ns.Labels, ns.Status)
			}
			select {
			case signals <- struct{}{}:
			default:
			}
		},
	})

	out := w.out
	var delayCh <-chan time.Time

	fmt.Println("hello informer")
	fmt.Println("hello informer")
	fmt.Println("hello informer")
	fmt.Println("hello informer")
	go w.informer.Run(w.catacomb.Dying())
	for {
		select {
		case <-w.catacomb.Dying():
			fmt.Println("dying")
			return w.catacomb.ErrDying()
		case <-signals:
			fmt.Println("signal")
			if delayCh == nil {
				delayCh = w.clock.After(sendDelay)
			}
		case <-delayCh:
			fmt.Println("delay")
			out = w.out
		case out <- struct{}{}:
			fmt.Println("out")
			logger.Debugf("fire notify watcher for %v", w.name)
			out = nil
			delayCh = nil
		}
	}
}

// Changes returns the event channel for this watcher.
func (w *kubernetesNotifyWatcher) Changes() watcher.NotifyChannel {
	return w.out
}

// Kill asks the watcher to stop without waiting for it do so.
func (w *kubernetesNotifyWatcher) Kill() {
	w.catacomb.Kill(nil)
}

// Wait waits for the watcher to die and returns any
// error encountered when it was running.
func (w *kubernetesNotifyWatcher) Wait() error {
	return w.catacomb.Wait()
}

type kubernetesStringsWatcher struct {
	clock    jujuclock.Clock
	catacomb catacomb.Catacomb

	out           chan []string
	name          string
	k8watcher     watch.Interface
	initialEvents []string
	filterFunc    k8sStringsWatcherFilterFunc
}

type k8sStringsWatcherFilterFunc func(evt watch.Event) (string, bool)

func newKubernetesStringsWatcher(wi watch.Interface, name string, clock jujuclock.Clock,
	initialEvents []string, filterFunc k8sStringsWatcherFilterFunc) (*kubernetesStringsWatcher, error) {
	w := &kubernetesStringsWatcher{
		clock:         clock,
		out:           make(chan []string),
		name:          name,
		k8watcher:     wi,
		initialEvents: initialEvents,
		filterFunc:    filterFunc,
	}
	err := catacomb.Invoke(catacomb.Plan{
		Site: &w.catacomb,
		Work: w.loop,
	})
	return w, err
}

func (w *kubernetesStringsWatcher) loop() error {
	defer close(w.out)
	defer w.k8watcher.Stop()

	select {
	case <-w.catacomb.Dying():
		return w.catacomb.ErrDying()
	case w.out <- w.initialEvents:
	}
	w.initialEvents = nil

	// Set out now so that initial event is sent.
	var out chan []string
	var delayCh <-chan time.Time
	var pendingEvents []string

	for {
		select {
		case <-w.catacomb.Dying():
			return w.catacomb.ErrDying()
		case evt, ok := <-w.k8watcher.ResultChan():
			// This can happen if the k8s API connection drops.
			if !ok {
				return errors.Errorf("k8s event watcher closed, restarting")
			}
			if evt.Type == watch.Error {
				return errors.Errorf("kubernetes watcher error: %v", k8serrors.FromObject(evt.Object))
			}
			logger.Tracef("received k8s event: %+v", evt.Type)
			if emittedEvent, ok := w.filterFunc(evt); ok {
				pendingEvents = append(pendingEvents, emittedEvent)
				if delayCh == nil {
					delayCh = w.clock.After(sendDelay)
				}
			}
		case <-delayCh:
			delayCh = nil
			out = w.out
		case out <- pendingEvents:
			out = nil
			pendingEvents = nil
		}
	}
}

// Changes returns the event channel for this watcher.
func (w *kubernetesStringsWatcher) Changes() watcher.StringsChannel {
	return w.out
}

// Kill asks the watcher to stop without waiting for it do so.
func (w *kubernetesStringsWatcher) Kill() {
	w.catacomb.Kill(nil)
}

// Wait waits for the watcher to die and returns any
// error encountered when it was running.
func (w *kubernetesStringsWatcher) Wait() error {
	return w.catacomb.Wait()
}
