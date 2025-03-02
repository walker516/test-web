import Header from "@/features/common/components/Header";

export default function UsersLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex h-screen">
      <div className="hidden lg:flex w-64 bg-gray-900 text-white flex-col">
        <div className="p-4 text-lg font-semibold">メニュー</div>
        <nav className="flex flex-col gap-2 p-4">
          <a href="/users" className="text-white hover:bg-gray-800 p-2 rounded">
            ユーザー一覧
          </a>
          <a href="/new" className="text-white hover:bg-gray-800 p-2 rounded">
            新規作成
          </a>
        </nav>
      </div>

      <div className="flex-1 flex flex-col">
        <Header />
        <main className="p-6 pt-20">{children}</main>
      </div>
    </div>
  );
}
