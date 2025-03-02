"use client";

import { useRouter } from "next/navigation";
import {
  Drawer,
  List,
  ListItem,
  ListItemText,
  IconButton,
  ListItemIcon,
} from "@mui/material";
import CloseTwoToneIcon from "@mui/icons-material/CloseTwoTone";
import GridViewIcon from "@mui/icons-material/GridView";
import AddIcon from "@mui/icons-material/Add";
import LogoutIcon from "@mui/icons-material/Logout";

export default function Sidebar({
  open,
  setOpen,
}: {
  open: boolean;
  setOpen: (state: boolean) => void;
}) {
  const router = useRouter();

  return (
    <Drawer
      anchor="left"
      open={open}
      onClose={() => setOpen(false)}
      className="lg:hidden"
    >
      <div className="w-64 h-full bg-gray-900 text-white flex flex-col">
        {/* 閉じるボタン */}
        <div className="p-4 flex justify-between items-center border-b border-gray-700">
          <span className="text-lg font-semibold">メニュー</span>
          <IconButton onClick={() => setOpen(false)} className="text-white">
            <CloseTwoToneIcon />
          </IconButton>
        </div>

        <List className="flex-1">
          <ListItem
            component="button"
            className="hover:bg-gray-800 transition duration-200"
            onClick={() => {
              router.push("/users");
              setOpen(false);
            }}
          >
            <ListItemIcon className="text-white">
              <GridViewIcon />
            </ListItemIcon>
            <ListItemText primary="ユーザー一覧" />
          </ListItem>

          <ListItem
            component="button"
            className="hover:bg-gray-800 transition duration-200"
            onClick={() => {
              router.push("/new");
              setOpen(false);
            }}
          >
            <ListItemIcon className="text-white">
              <AddIcon />
            </ListItemIcon>
            <ListItemText primary="新規作成" />
          </ListItem>
        </List>

        {/* ログアウトボタン */}
        <div className="p-4 border-t border-gray-700">
          <button
            className="w-full flex items-center justify-center bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-md transition duration-200"
            onClick={() => alert("ログアウト")}
          >
            <LogoutIcon className="mr-2" />
            Logout
          </button>
        </div>
      </div>
    </Drawer>
  );
}
