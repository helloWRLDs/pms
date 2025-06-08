-- +goose Up
-- +goose StatementBegin

-- Participant role (basic access)
INSERT INTO "Role" (name, permissions) 
VALUES (
    'participant',
    ARRAY[
        'project:read',
        'task:read',
        'task:write',
        'sprint:read'
    ]
) ON CONFLICT (name) DO NOTHING;

-- Project Manager role (includes participant permissions plus more)
INSERT INTO "Role" (name, permissions) 
VALUES (
    'project_manager',
    ARRAY[
        'project:read',
        'project:write',
        'project:invite',
        'task:read',
        'task:write',
        'task:invite',
        'user:read',
        'user:invite',
        'sprint:read',
        'sprint:write',
        'sprint:invite'
    ]
) ON CONFLICT (name) DO NOTHING;

-- CEO role (includes all permissions)
INSERT INTO "Role" (name, permissions) 
VALUES (
    'ceo',
    ARRAY[
        'org:read',
        'org:write',
        'user:read',
        'user:write',
        'user:delete',
        'user:invite',
        'project:read',
        'project:write',
        'project:delete',
        'project:invite',
        'task:read',
        'task:write',
        'task:delete',
        'task:invite',
        'role:read',
        'role:write',
        'role:delete',
        'role:invite',
        'sprint:read',
        'sprint:write',
        'sprint:delete',
        'sprint:invite'
    ]
) ON CONFLICT (name) DO NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "Role"
WHERE
    name IN (
        'participant',
        'project_manager',
        'ceo'
    );
-- +goose StatementEnd