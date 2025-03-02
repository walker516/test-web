"use client";

import { useUsers } from "../hooks/useUsers";
import { AgGridReact } from "ag-grid-react";
import {
  ClientSideRowModelModule,
  ModuleRegistry,
  ColDef,
} from "ag-grid-community";
import { Card, CardContent, Typography, CircularProgress } from "@mui/material";
import { useRouter } from "next/navigation";

ModuleRegistry.registerModules([ClientSideRowModelModule]);

export default function UserTable() {
  const { users, isLoading } = useUsers();
  const router = useRouter();

  const columns: ColDef[] = [
    { field: "id", headerName: "ID", sortable: true, filter: true },
    {
      field: "name",
      headerName: "名前",
      sortable: true,
      filter: "agTextColumnFilter",
    },
    {
      field: "email",
      headerName: "メール",
      sortable: true,
      filter: "agTextColumnFilter",
    },
    {
      field: "role",
      headerName: "権限",
      sortable: true,
      filter: "agSetColumnFilter",
    },
  ];

  return (
    <Card className="shadow-lg rounded-xl">
      <CardContent>
        <Typography variant="h6" className="mb-4">
          ユーザー一覧
        </Typography>

        {isLoading ? (
          <div className="flex justify-center py-6">
            <CircularProgress />
          </div>
        ) : (
          <div className="ag-theme-alpine h-[500px] w-full">
            <AgGridReact
              rowData={users}
              columnDefs={columns}
              domLayout="autoHeight"
              rowSelection="single"
              onRowClicked={(params) => router.push(`/users/${params.data.id}`)}
              rowClass="cursor-pointer hover:bg-gray-100"
            />
          </div>
        )}
      </CardContent>
    </Card>
  );
}
