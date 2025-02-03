// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)


// ChangeLogTriggersForModelSecretBackend generates the triggers for the
// model_secret_backend table.
func ChangeLogTriggersForModelSecretBackend(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for ModelSecretBackend
INSERT INTO change_log_namespace VALUES (%[2]d, 'model_secret_backend', 'ModelSecretBackend changes based on %[1]s');

-- insert trigger for ModelSecretBackend
CREATE TRIGGER trg_log_model_secret_backend_insert
AFTER INSERT ON model_secret_backend FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for ModelSecretBackend
CREATE TRIGGER trg_log_model_secret_backend_update
AFTER UPDATE ON model_secret_backend FOR EACH ROW
WHEN 
	NEW.model_uuid != OLD.model_uuid OR
	NEW.secret_backend_uuid != OLD.secret_backend_uuid 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for ModelSecretBackend
CREATE TRIGGER trg_log_model_secret_backend_delete
AFTER DELETE ON model_secret_backend FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForSecretBackendRotation generates the triggers for the
// secret_backend_rotation table.
func ChangeLogTriggersForSecretBackendRotation(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for SecretBackendRotation
INSERT INTO change_log_namespace VALUES (%[2]d, 'secret_backend_rotation', 'SecretBackendRotation changes based on %[1]s');

-- insert trigger for SecretBackendRotation
CREATE TRIGGER trg_log_secret_backend_rotation_insert
AFTER INSERT ON secret_backend_rotation FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for SecretBackendRotation
CREATE TRIGGER trg_log_secret_backend_rotation_update
AFTER UPDATE ON secret_backend_rotation FOR EACH ROW
WHEN 
	NEW.backend_uuid != OLD.backend_uuid OR
	NEW.next_rotation_time != OLD.next_rotation_time 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for SecretBackendRotation
CREATE TRIGGER trg_log_secret_backend_rotation_delete
AFTER DELETE ON secret_backend_rotation FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

