<?php

class UserGateway
{
    private PDO $conn;

    public function __construct(Database $database)
    {
        $this->conn = $database->getConnection();
    }

    public function getAll(): array
    {
        $sql = "SELECT * 
            FROM user";

        $stmt = $this->conn->query($sql);

        $data = [];

        while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
            $data[] = $row;
        }

        return $data;

    }

    public function create(array $data): string
    {
        $sql = "INSERT INTO user (pseudo, email, passwd, id_imagepp, theme) 
                VALUES (:pseudo, :email, :passwd, :id_imagepp, :theme)";

        $stmt = $this->conn->prepare($sql);

        $options = [
            'cost' => 14,
        ];
        $data["passwd"] = password_hash(htmlspecialchars($data["passwd"]), PASSWORD_DEFAULT, $options);

        $stmt->bindValue(":pseudo", htmlspecialchars($data["pseudo"]), PDO::PARAM_STR);
        $stmt->bindValue(":email", htmlspecialchars($data["email"]), PDO::PARAM_STR);
        $stmt->bindValue(":passwd", $data["passwd"], PDO::PARAM_STR);
        $stmt->bindValue(":id_imagepp", htmlspecialchars($data["id_imagepp"]), PDO::PARAM_INT);
        $stmt->bindValue(":theme", htmlspecialchars($data["theme"]), PDO::PARAM_STR);

        $stmt->execute();

        return $this->conn->lastInsertId();

    }

    public function get(string $id): array|false
    {
        $sql = "SELECT *
                FROM user
                WHERE id_user = :id_user";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_user", htmlspecialchars($id), PDO::PARAM_INT);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

    public function getByPseudo(string $pseudo): array|false
    {
        $sql = "SELECT *
                FROM user
                WHERE pseudo = :pseudo";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":pseudo", htmlspecialchars($pseudo), PDO::PARAM_STR);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

    public function getByEmail(string $pseudo): array|false
    {
        $sql = "SELECT *
                FROM user
                WHERE email = :email";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":email", htmlspecialchars($pseudo), PDO::PARAM_STR);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

    public function update(array $current, array $new): int
    {
        $sql = "UPDATE user
                SET pseudo = :pseudo, email = :email, passwd =:passwd, id_imagepp= :id_imagepp, theme = :theme
                WHERE id_user = :id_user";

        $stmt = $this->conn->prepare($sql);

        $options = [
            'cost' => 14,
        ];

        if (isset($new["pseudo"])) {
            $new["pseudo"] = htmlspecialchars($new["pseudo"]);
        }
        if (isset($new["email"])) {
            $new["email"] = htmlspecialchars($new["email"]);
        }
        if (isset($new["passwd"])) {
            $new["passwd"] = password_hash(htmlspecialchars($new["passwd"]), PASSWORD_DEFAULT, $options);
        }
        if (isset($new["id_imagepp"])) {
            $new["id_imagepp"] = htmlspecialchars($new["id_imagepp"]);
        }
        if (isset($new["theme"])) {
            $new["theme"] = htmlspecialchars($new["theme"]);
        }

        $stmt->bindValue(
            ":pseudo", $new["pseudo"] ?? $current["pseudo"],
            PDO::PARAM_STR
        );
        $stmt->bindValue(
            ":email", $new["email"] ?? $current["email"],
            PDO::PARAM_STR
        );
        $stmt->bindValue(
            ":passwd", $new["passwd"] ?? $current["passwd"],
            PDO::PARAM_STR
        );
        $stmt->bindValue(
            ":id_imagepp", $new["id_imagepp"] ?? $current["id_imagepp"],
            PDO::PARAM_INT
        );
        $stmt->bindValue(
            ":theme", $new["theme"] ?? $current["theme"],
            PDO::PARAM_STR
        );

        $stmt->bindValue(":id_user", $current["id_user"], PDO::PARAM_INT);

        $stmt->execute();

        return $stmt->rowCount();
    }

    public function delete(string $id): int
    {
        $sql = "DELETE FROM user
                WHERE id_user = :id_user";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_user", htmlspecialchars($id), PDO::PARAM_INT);

        $stmt->execute();

        return $stmt->rowCount();
    }

}