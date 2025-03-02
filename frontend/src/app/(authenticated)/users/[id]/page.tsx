"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import { UserRepository } from "@/domain/user/Repository";
import { User } from "@/domain/user/User";
import {
  Card,
  CardContent,
  Typography,
  Button,
  CircularProgress,
} from "@mui/material";

export default function UserDetailPage() {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();
  const params = useParams();

  useEffect(() => {
    if (!params.id || typeof params.id !== "string") return;

    setIsLoading(true);
    UserRepository.getById(Number(params.id))
      .then(setUser)
      .catch(() => setUser(null))
      .finally(() => setIsLoading(false));
  }, [params.id]);

  if (isLoading) {
    return (
      <div className="flex justify-center py-6">
        <CircularProgress />
      </div>
    );
  }

  if (!user) {
    return (
      <Typography variant="body1" className="text-center">
        ユーザーが見つかりません
      </Typography>
    );
  }

  return (
    <Card className="max-w-md mx-auto shadow-lg rounded-xl">
      <CardContent>
        <Typography variant="h6" className="mb-4">
          ユーザー詳細
        </Typography>
        <Typography variant="body1">
          <strong>名前:</strong> {user.name}
        </Typography>
        <Typography variant="body1">
          <strong>Email:</strong> {user.email}
        </Typography>
        <Typography variant="body1">
          <strong>権限:</strong> {user.role}
        </Typography>
        <Button
          variant="contained"
          color="primary"
          className="mt-4 w-full"
          onClick={() => router.push(`/users/${user.id}/edit`)}
        >
          編集
        </Button>
        <Button variant="outlined" onClick={() => router.push("/users")}>
          戻る
        </Button>
      </CardContent>
    </Card>
  );
}
