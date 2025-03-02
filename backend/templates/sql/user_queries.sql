-- START: GetByID
SELECT id, name, email, role, created_at
FROM users
WHERE id = :Id;
-- END

-- START: GetAll
SELECT id, name, email, role, created_at
FROM users;
-- END

-- START: Create
INSERT INTO users (name, email, role, created_at)
VALUES (:Name, :Email, :Role, NOW());
-- END

-- START: UpdateUser
UPDATE users
SET name = :Name, email = :Email, role = :Role
WHERE id = :Id;
-- END

-- START: DeleteUser
DELETE FROM users WHERE id = :Id;
-- END
