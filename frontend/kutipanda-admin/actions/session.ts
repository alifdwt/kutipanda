import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { SignJWT, jwtVerify } from "jose";
import "server-only";

const JWTSecret = process.env.JWT_SECRET;
const key = new TextEncoder().encode(JWTSecret);

export async function decrypt(session: string | undefined = "") {
  try {
    const { payload } = await jwtVerify(session, key, {
      algorithms: ["HS256"],
    });
    return payload;
  } catch (error) {
    return null;
  }
}

export async function createSession(
  access_token: string,
  refresh_token: string
) {
  const expiresAt = new Date(Date.now() + 60 * 60 * 1000);

  cookies().set("access_token", access_token, {
    httpOnly: true,
    secure: true,
    expires: expiresAt,
    sameSite: "lax",
    path: "/",
  });

  cookies().set("refresh_token", refresh_token, {
    httpOnly: true,
    secure: true,
    expires: expiresAt,
    sameSite: "lax",
    path: "/",
  });

  redirect("/dashboard");
}

export async function verifySession() {
  const cookie = cookies().get("access_token")?.value;
  const session = await decrypt(cookie);

  if (!session?.sub) {
    redirect("/login");
  }

  return { isAuth: true, userId: Number(session.sub) };
}

export function deleteSession() {
  cookies().delete("access_token");
  cookies().delete("refresh_token");

  redirect("/login");
}
