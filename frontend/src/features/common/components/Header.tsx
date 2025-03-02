"use client";

import { useState } from "react";
import { AppBar, Toolbar, Typography, IconButton, Box } from "@mui/material";
import MenuTwoToneIcon from "@mui/icons-material/MenuTwoTone";
import LogoutIcon from "@mui/icons-material/Logout";
import Sidebar from "./Sidebar";

export default function Header() {
  const [sidebarOpen, setSidebarOpen] = useState(false);

  return (
    <>
      <AppBar position="fixed" className="bg-blue-600 shadow-md">
        <Toolbar className="flex justify-between">
          {/* ハンバーガーメニュー (モバイル用) */}
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            className="lg:hidden"
            onClick={() => setSidebarOpen(!sidebarOpen)}
          >
            <MenuTwoToneIcon />
          </IconButton>

          <Typography variant="h6" className="text-white">
            ユーザー管理システム
          </Typography>

          {/* ログアウトボタン */}
          <Box>
            <IconButton
              className="text-white hover:bg-red-600 transition duration-200"
              onClick={() => alert("ログアウト")}
            >
              <LogoutIcon />
            </IconButton>
          </Box>
        </Toolbar>
      </AppBar>

      {/* Sidebar (ドロワー) */}
      <Sidebar open={sidebarOpen} setOpen={setSidebarOpen} />
    </>
  );
}
