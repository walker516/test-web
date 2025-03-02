import { useEffect, useState } from "react";
import { UserRepository } from "@/domain/user/Repository";
import { User } from "@/domain/user/User";

export function useUsers() {
  const [users, setUsers] = useState<User[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    UserRepository.getAll()
      .then((data) => {
        console.log("Fetched users:", data);
        setUsers(data);
      })
      .catch((error) => {
        console.error("Error fetching users:", error);
      })
      .finally(() => setIsLoading(false));
  }, []);

  return { users, isLoading };
}
