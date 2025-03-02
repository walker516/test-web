"use client";

import UserTable from "@/features/users/components/UserTable";
import { Button } from "@mui/material";
import { useRouter } from "next/navigation";

export default function UsersPage() {
  const router = useRouter();
  return (
    <div className="p-4">
      <h2 className="text-lg font-semibold mb-4">ユーザー一覧</h2>
      <Button
        variant="contained"
        color="primary"
        onClick={() => router.push("/new")}
      >
        新規作成
      </Button>
      <UserTable />
    </div>
  );
}
