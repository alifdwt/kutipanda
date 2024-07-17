import { Metadata } from "next";
import Link from "next/link";
import { LoginForm } from "./_components/login-form";

export const metadata: Metadata = {
  title: "Login",
  description: "Login to your account",
  keywords: ["kutipanda", "login", "kutipanda admin", "kutipanda admin login"],
};

export default function LoginPage() {
  return (
    <div className="flex flex-col p-4 lg:w-1/3">
      <div className="text-center">
        <h1 className="text-3xl font-bold">Masuk</h1>
        <p className="text-gray-500">
          Selamat datang kembali! <br /> Silakan masukkan email dan kata sandi
          Anda
        </p>
      </div>
      <div className="mt-6">
        <LoginForm />
      </div>
      <div className="mt-4 text-center text-sm">
        <span className="text-gray-500">Belum punya akun? </span>
        <Link className="underline" href="/signup">
          Sign up
        </Link>
      </div>
    </div>
  );
}
