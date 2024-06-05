import { cookies } from "next/headers";
import Link from "next/link";
import UserButton from "./UserButton";
import { Button } from "../ui/button";

const Header = () => {
  const cookie = cookies();
  const token = cookie.get("access_token")?.value;

  return (
    <header className="w-full border-b">
      <div className="wrapper flex items-center justify-between">
        <Link href="/" className="w-36">
          <h1 className="text-3xl font-bold">KUTIPANDA</h1>
        </Link>

        <nav className="md:flex-between hidden w-full max-w-xs">
          {/* TODO */}
        </nav>

        <div className="flex w-32 justify-end gap-3">
          {token ? (
            <UserButton access_token={token} />
          ) : (
            <Button asChild>
              <Link href="/sign-in">Sign In</Link>
            </Button>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
