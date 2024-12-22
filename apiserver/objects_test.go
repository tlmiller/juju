// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package apiserver

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	jc "github.com/juju/testing/checkers"
	gomock "go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/apiserverhttp"
	"github.com/juju/juju/domain/application/architecture"
	applicationcharm "github.com/juju/juju/domain/application/charm"
	applicationerrors "github.com/juju/juju/domain/application/errors"
	"github.com/juju/juju/internal/testing"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
)

type objectsCharmHandlerSuite struct {
	applicationsServiceGetter *MockApplicationServiceGetter
	applicationsService       *MockApplicationService

	// These will move to the model service.
	stateGetter *MockStateGetter
	modelState  *MockModelState
	model       *MockModel

	mux *apiserverhttp.Mux
	srv *httptest.Server
}

var _ = gc.Suite(&objectsCharmHandlerSuite{})

func (s *objectsCharmHandlerSuite) SetUpTest(c *gc.C) {
	s.mux = apiserverhttp.NewMux()
	s.srv = httptest.NewServer(s.mux)
}

func (s *objectsCharmHandlerSuite) TearDownTest(c *gc.C) {
	s.srv.Close()
}

func (s *objectsCharmHandlerSuite) TestServeMethodNotSupported(c *gc.C) {
	defer s.setupMocks(c).Finish()

	handlers := &objectsCharmHTTPHandler{
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	// This is a bit pathological, but we want to make sure that the handler
	// logic only actions on POST requests.
	s.mux.AddHandler("POST", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("POST", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "0abcdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm-%s", s.srv.URL, modelUUID, hashPrefix)
	resp, err := http.Post(url, "application/octet-stream", strings.NewReader("charm-content"))
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(resp.StatusCode, gc.Equals, http.StatusNotImplemented)
}

func (s *objectsCharmHandlerSuite) TestServeGet(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectApplicationService()

	s.applicationsService.EXPECT().GetCharmArchiveBySHA256Prefix(gomock.Any(), "01abcdef").Return(io.NopCloser(strings.NewReader("charm-content")), nil)

	handlers := &objectsCharmHTTPHandler{
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	s.mux.AddHandler("GET", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("GET", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "01abcdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm-%s", s.srv.URL, modelUUID, hashPrefix)
	resp, err := http.Get(url)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(resp.StatusCode, gc.Equals, http.StatusOK)
	body, err := io.ReadAll(resp.Body)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(string(body), gc.Equals, "charm-content")
}

func (s *objectsCharmHandlerSuite) TestServeGetNotFound(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectApplicationService()

	s.applicationsService.EXPECT().GetCharmArchiveBySHA256Prefix(gomock.Any(), "01abcdef").Return(nil, applicationerrors.CharmNotFound)

	handlers := &objectsCharmHTTPHandler{
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	s.mux.AddHandler("GET", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("GET", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "01abcdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm-%s", s.srv.URL, modelUUID, hashPrefix)
	resp, err := http.Get(url)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(resp.StatusCode, gc.Equals, http.StatusNotFound)
}

func (s *objectsCharmHandlerSuite) TestServePutIncorrectEncoding(c *gc.C) {
	defer s.setupMocks(c).Finish()

	handlers := &objectsCharmHTTPHandler{
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	s.mux.AddHandler("PUT", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("PUT", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "01abcdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm-%s", s.srv.URL, modelUUID, hashPrefix)
	req, err := http.NewRequest("PUT", url, strings.NewReader("charm-content"))
	c.Assert(err, jc.ErrorIsNil)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(resp.StatusCode, gc.Equals, http.StatusBadRequest)
}

func (s *objectsCharmHandlerSuite) TestServePutNoJujuCharmURL(c *gc.C) {
	defer s.setupMocks(c).Finish()

	handlers := &objectsCharmHTTPHandler{
		stateGetter:              s.stateGetter,
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	s.expectApplicationService()
	s.expectModelState()

	s.mux.AddHandler("PUT", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("PUT", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "01abcdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm-%s", s.srv.URL, modelUUID, hashPrefix)
	req, err := http.NewRequest("PUT", url, strings.NewReader("charm-content"))
	c.Assert(err, jc.ErrorIsNil)

	req.Header.Set("Content-Type", "application/zip")

	resp, err := http.DefaultClient.Do(req)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(resp.StatusCode, gc.Equals, http.StatusBadRequest)
}

func (s *objectsCharmHandlerSuite) TestServePutInvalidSHA256Prefix(c *gc.C) {
	defer s.setupMocks(c).Finish()

	handlers := &objectsCharmHTTPHandler{
		stateGetter:              s.stateGetter,
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	s.expectApplicationService()
	s.expectModelState()

	s.mux.AddHandler("PUT", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("PUT", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "cdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm-%s", s.srv.URL, modelUUID, hashPrefix)
	req, err := http.NewRequest("PUT", url, strings.NewReader("charm-content"))
	c.Assert(err, jc.ErrorIsNil)

	req.Header.Set("Content-Type", "application/zip")
	req.Header.Set(params.JujuCharmURLHeader, "ch:testcharm-1")

	resp, err := http.DefaultClient.Do(req)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(resp.StatusCode, gc.Equals, http.StatusBadRequest)
}

func (s *objectsCharmHandlerSuite) TestServePutInvalidCharmURL(c *gc.C) {
	defer s.setupMocks(c).Finish()

	handlers := &objectsCharmHTTPHandler{
		stateGetter:              s.stateGetter,
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	s.expectApplicationService()
	s.expectModelState()

	s.mux.AddHandler("PUT", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("PUT", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "01abcdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm_%s", s.srv.URL, modelUUID, hashPrefix)
	req, err := http.NewRequest("PUT", url, strings.NewReader("charm-content"))
	c.Assert(err, jc.ErrorIsNil)

	req.Header.Set("Content-Type", "application/zip")
	req.Header.Set(params.JujuCharmURLHeader, "ch:testcharm-1")

	resp, err := http.DefaultClient.Do(req)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(resp.StatusCode, gc.Equals, http.StatusBadRequest)
}

func (s *objectsCharmHandlerSuite) TestServePut(c *gc.C) {
	defer s.setupMocks(c).Finish()

	handlers := &objectsCharmHTTPHandler{
		stateGetter:              s.stateGetter,
		applicationServiceGetter: s.applicationsServiceGetter,
	}

	s.expectApplicationService()
	s.expectModelState()
	s.expectModel()

	s.mux.AddHandler("PUT", charmsObjectsRoutePrefix, handlers)
	defer s.mux.RemoveHandler("PUT", charmsObjectsRoutePrefix)

	modelUUID := testing.ModelTag.Id()
	hashPrefix := "01abcdef"

	url := fmt.Sprintf("%s/model-%s/charms/testcharm-%s", s.srv.URL, modelUUID, hashPrefix)
	req, err := http.NewRequest("PUT", url, strings.NewReader("charm-content"))
	c.Assert(err, jc.ErrorIsNil)

	req.Header.Set("Content-Type", "application/zip")
	req.Header.Set(params.JujuCharmURLHeader, "ch:testcharm-1")

	s.applicationsService.EXPECT().ResolveUploadCharm(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, args applicationcharm.ResolveUploadCharm) (applicationcharm.CharmLocator, error) {
		c.Check(args.Name, gc.Equals, "testcharm")
		c.Check(args.Revision, gc.Equals, 1)
		c.Check(args.Architecture, gc.Equals, "")

		return applicationcharm.CharmLocator{
			Name:         "testcharm",
			Revision:     2,
			Source:       applicationcharm.CharmHubSource,
			Architecture: architecture.AMD64,
		}, nil

	})

	resp, err := http.DefaultClient.Do(req)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(resp.StatusCode, gc.Equals, http.StatusOK)
	c.Check(resp.Header.Get(params.JujuCharmURLHeader), gc.Equals, "ch:amd64/testcharm-2")
}

func (s *objectsCharmHandlerSuite) expectApplicationService() {
	s.applicationsServiceGetter.EXPECT().Application(gomock.Any()).Return(s.applicationsService, nil)
}

func (s *objectsCharmHandlerSuite) expectModelState() {
	s.stateGetter.EXPECT().GetState(gomock.Any()).Return(s.modelState, nil)
	s.modelState.EXPECT().Release().Return(true)
}

func (s *objectsCharmHandlerSuite) expectModel() {
	s.modelState.EXPECT().Model().Return(s.model, nil)
	s.model.EXPECT().MigrationMode().Return(state.MigrationModeNone)
}

func (s *objectsCharmHandlerSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.applicationsServiceGetter = NewMockApplicationServiceGetter(ctrl)
	s.applicationsService = NewMockApplicationService(ctrl)

	// These should be on the model service!
	s.stateGetter = NewMockStateGetter(ctrl)
	s.modelState = NewMockModelState(ctrl)
	s.model = NewMockModel(ctrl)

	return ctrl
}
