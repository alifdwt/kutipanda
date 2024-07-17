import "server-only";
import { verifySession } from "./session";
import { cache } from "react";
import { BASE_URL } from "@/lib/api";
import { cookies } from "next/headers";

export const getUser = cache(async () => {
  const cookie = cookies().get("access_token")?.value;
  const session = await verifySession();
  if (!session) return null;

  try {
    const data = await fetch(`${BASE_URL}/user/me`, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${cookie}`,
      },
    }).then((res) => res.json());

    return data.data;
  } catch (error) {
    console.error(error);
    return null;
  }
});
