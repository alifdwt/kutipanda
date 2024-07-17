"use client";

import { login } from "@/actions/auth";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Link from "next/link";
import { useFormState, useFormStatus } from "react-dom";

export function LoginForm() {
  const [state, action] = useFormState(login, undefined);
  const { pending } = useFormStatus();

  return (
    <form action={action}>
      <div className="flex flex-col gap-4">
        <div>
          <Label htmlFor="email">Email</Label>
          <Input
            id="email"
            type="email"
            name="email"
            placeholder="email@mail.com"
          />
          {state?.errors?.email && (
            <p className="text-sm text-red-500">{state.errors.email}</p>
          )}
        </div>
        <div>
          <div className="flex items-center justify-between">
            <Label htmlFor="password">Password</Label>
            <Link className="text-sm underline" href="#">
              Forgot Password?
            </Link>
          </div>
          <Input
            id="password"
            type="password"
            name="password"
            placeholder="password"
          />
          {state?.errors?.password && (
            <p className="text-sm text-red-500">{state.errors.password}</p>
          )}
        </div>
        {state?.message && (
          <p className="text-sm text-red-500">{state.message}</p>
        )}
        <Button aria-disabled={pending} type="submit" className="mt-4 w-full">
          {pending ? "Logging in..." : "Login"}
        </Button>
      </div>
    </form>
  );
}
