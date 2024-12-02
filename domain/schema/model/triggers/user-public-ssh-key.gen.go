// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)

// ChangeLogTriggersForUserPublicSshKey generates the triggers for the
// user_public_ssh_key table.
func ChangeLogTriggersForUserPublicSshKey(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for UserPublicSshKey
INSERT INTO change_log_namespace VALUES (%[2]d, 'user_public_ssh_key', 'UserPublicSshKey changes based on %[1]s');

-- insert trigger for UserPublicSshKey
CREATE TRIGGER trg_log_user_public_ssh_key_insert
AFTER INSERT ON user_public_ssh_key FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for UserPublicSshKey
CREATE TRIGGER trg_log_user_public_ssh_key_update
AFTER UPDATE ON user_public_ssh_key FOR EACH ROW
WHEN 
	(NEW.id != OLD.id OR (NEW.id IS NOT NULL AND OLD.id IS NULL) OR (NEW.id IS NULL AND OLD.id IS NOT NULL)) OR
	NEW.comment != OLD.comment OR
	NEW.fingerprint_hash_algorithm_id != OLD.fingerprint_hash_algorithm_id OR
	NEW.fingerprint != OLD.fingerprint OR
	NEW.public_key != OLD.public_key OR
	NEW.user_id != OLD.user_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for UserPublicSshKey
CREATE TRIGGER trg_log_user_public_ssh_key_delete
AFTER DELETE ON user_public_ssh_key FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}
