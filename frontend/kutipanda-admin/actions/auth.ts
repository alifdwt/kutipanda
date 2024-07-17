"use server";

import { BASE_URL } from "@/lib/api";
import { FormState, LoginFormSchema } from "@/lib/validation/login";
import { createSession, deleteSession } from "./session";

export async function login(
  state: FormState,
  formData: FormData
): Promise<FormState> {
  const validateFields = LoginFormSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  });
  const errorMessage = { message: "Invalid login credentials" };

  if (!validateFields.success) {
    return {
      errors: validateFields.error.flatten().fieldErrors,
    };
  }

  const response = await fetch(`${BASE_URL}/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(validateFields.data),
  })
    .then((res) => res.json())
    .catch(() => errorMessage);

  if (response.statusCode !== 200) {
    return {
      message: response.message,
    };
  }

  await createSession(response.data.access_token, response.data.refresh_token);
}

export async function logout() {
  deleteSession();
}
