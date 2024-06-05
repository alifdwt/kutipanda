import LoginForm from "./_components/LoginForm";

export default function SignInPage() {
  return (
    <div className="bg-gradient-to-br from-purple-700 to-pink-500 min-h-screen flex flex-col justify-center items-center">
      <div className="bg-white rounded-lg shadow-lg p-8 max-w-md">
        <h1 className="text-4xl font-bold text-center text-purple-700 mb-8">
          Silakan Masuk
        </h1>
        <LoginForm />
      </div>
    </div>
  );
}
