"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import UserForm from "@/features/users/components/UserForm";
import { UserRepository } from "@/domain/user/Repository";
import { User } from "@/domain/user/User";

export default function EditUserPage() {
  const [user, setUser] = useState<User | null>(null);
  const router = useRouter();
  const params = useParams(); // useParams を利用

  useEffect(() => {
    if (!params || typeof params.id !== "string") return;

    UserRepository.getById(Number(params.id))
      .then(setUser)
      .catch(() => setUser(null));
  }, [params]);

  if (!user) return <p>ユーザーが見つかりません</p>;

  return (
    <div>
      <h2>ユーザー編集</h2>
      <UserForm user={user} />
    </div>
  );
}
