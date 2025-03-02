import axiosClient from "@/utils/axiosClient";
import { User } from "./User";

export class UserRepository {
  static async getAll(): Promise<User[]> {
    const res = await axiosClient.get("/api/user/v1/users");
    return res.data;
  }

  static async getById(id: number): Promise<User> {
    const res = await axiosClient.get(`/api/user/v1/users/${id}`);
    return res.data;
  }

  static async create(user: Omit<User, "id">) {
    return await axiosClient.post("/api/user/v1/users", user);
  }

  static async update(id: number, user: Partial<User>) {
    return await axiosClient.put(`/api/user/v1/users/${id}`, user);
  }

  static async delete(id: number) {
    return await axiosClient.delete(`/api/user/v1/users/${id}`);
  }
}