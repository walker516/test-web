import { Button, MenuItem, Select, TextField, Stack } from "@mui/material";
import { useUserForm } from "../hooks/useUserForm";

export default function UserForm({ user }: { user?: any }) {
  const { form, handleChange, handleSubmit } = useUserForm(user);

  return (
    <form onSubmit={handleSubmit} className="max-w-md mx-auto">
      <Stack spacing={3}>
        <TextField
          label="名前"
          name="name"
          value={form.name}
          onChange={handleChange}
          fullWidth
          required
        />
        <TextField
          label="メール"
          name="email"
          type="email"
          value={form.email}
          onChange={handleChange}
          fullWidth
          required
        />
        <Select
          name="role"
          value={form.role}
          onChange={handleChange}
          fullWidth
          displayEmpty
        >
          <MenuItem value="" disabled>
            権限を選択
          </MenuItem>
          <MenuItem value="user">User</MenuItem>
          <MenuItem value="admin">Admin</MenuItem>
        </Select>
        <Button type="submit" variant="contained" color="primary" fullWidth>
          {user ? "更新" : "作成"}
        </Button>
      </Stack>
    </form>
  );
}
