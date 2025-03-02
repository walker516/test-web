import { useState } from "react";
import { useRouter } from "next/navigation";
import { UserRepository } from "@/domain/user/Repository";
import { SelectChangeEvent } from "@mui/material";

export function useUserForm(initialUser?: any) {
  const [form, setForm] = useState(
    initialUser || { name: "", email: "", role: "user" }
  );
  const router = useRouter();

  const handleChange = (
    e:
      | React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
      | SelectChangeEvent
  ) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (initialUser) {
      await UserRepository.update(initialUser.id, form);
    } else {
      await UserRepository.create(form);
    }
    router.push("/users");
  };

  return { form, handleChange, handleSubmit };
}
