import { siteConfig } from "@/app";
import { MainMenu } from "./MainMenu";
import { MobileMenu } from "./MobileMenu";
import { dashboardConfig } from "@/constants/nav-items";
import { ThemesGeneralSwitcher } from "../Switchers/ThemeGeneralSwitcher";
import { DiscordLogoIcon, GitHubLogoIcon } from "@radix-ui/react-icons";
import Link from "next/link";
import { buttonVariants } from "../ui/button";
import { cn } from "@/lib/utils";

export async function SiteHeader() {
  return (
    <header
      className={`
        container sticky top-0 z-40 mb-2 flex h-16 w-full justify-between
        border-b bg-background/90 backdrop-blur-sm duration-slow items-center
        animate-in fade-in slide-in-from-top-full
          `}
    >
      <MainMenu items={siteConfig.mainNav} />
      <MobileMenu
        MainMenuItems={siteConfig.mainNav}
        sidebarNavItems={dashboardConfig.sidebarNav}
      />
      <nav className="flex flex-1 items-center justify-end space-x-2">
        <ThemeDropdown />

        <div className="flex space-x-2">
          <GithubLink />
          <DiscordLink />
        </div>
      </nav>
    </header>
  );
}

function ThemeDropdown() {
  if (!siteConfig.themeToggleEnabled) {
    return null;
  }

  return <ThemesGeneralSwitcher />;
}

function DiscordLink() {
  return (
    <Link
      className={cn(
        buttonVariants({
          size: "icon",
          variant: "outline",
        }),
        `
            hidden
  
            sm:flex
          `
      )}
      href="https://discord.gg/TK2SpfXfTB"
      rel="noreferrer noopener"
      target="_blank"
    >
      <DiscordLogoIcon className="size-4" />
    </Link>
  );
}

function GithubLink() {
  return (
    <Link
      className={cn(
        buttonVariants({
          size: "icon",
          variant: "outline",
        })
      )}
      href="https://github.com/blefnk/relivator-nextjs-template"
      rel="noreferrer noopener"
      target="_blank"
    >
      <GitHubLogoIcon className="size-4" />
    </Link>
  );
}
