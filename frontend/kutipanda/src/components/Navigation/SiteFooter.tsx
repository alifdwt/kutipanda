import { cn } from "@/lib/utils";
import Link from "next/link";
import { ReactNode } from "react";
import { buttonVariants } from "../ui/button";
import { Shell } from "../Wrappers/ShellVariants";
import { EarthIcon, MusicIcon } from "lucide-react";
import { siteConfig } from "@/app";
import { GitHubLogoIcon } from "@radix-ui/react-icons";
import { ThemesGeneralSwitcher } from "../Switchers/ThemeGeneralSwitcher";
import { env } from "@/env";

type SocialIconLinkProps = {
  children: ReactNode;
  href: string;
  label: string;
};

const SocialIconLink = ({ children, href, label }: SocialIconLinkProps) => (
  <Link
    aria-label={label}
    className={cn(
      buttonVariants({
        size: "icon",
        variant: "outline",
      })
    )}
    href={href}
    rel="noreferrer noopener"
    target="_blank"
  >
    {children}
  </Link>
);

export function SiteFooter() {
  return (
    <footer className="items-center border-t bg-background py-14">
      <Shell
        aria-labelledby="footer-content-heading"
        as="nav"
        className={`
          flex max-w-screen-xl flex-col gap-10

          xl:flex-row
        `}
        id="footer-content"
      >
        <section
          aria-labelledby="footer-newsletter-with-socials-and-copyright"
          className="flex flex-col items-center space-y-8"
          id="footer-newsletter-with-socials-and-copyright"
        >
          <div>
            <Link
              className="mb-4 flex w-fit items-center space-x-4 font-medium"
              href="/"
            >
              <EarthIcon aria-hidden="true" className="size-6" />
              <span>{siteConfig.appNameDesc}</span>
              <span className="sr-only">
                Website logo and full name, with link to home page
              </span>
            </Link>
          </div>
          <div className="flex space-x-4">
            <div className="flex items-center space-x-1">
              {/* <SocialIconLink href={siteConfig.socialLinks.} label="Discord">
                <DiscordLogoIcon aria-hidden="true" className="size-4" />
              </SocialIconLink> */}
              <SocialIconLink
                href={siteConfig.socialLinks.github}
                label="GitHub"
              >
                <GitHubLogoIcon aria-hidden="true" className="size-4" />
              </SocialIconLink>
              <SocialIconLink
                href="https://youtube.com/@alifdwt"
                label="YouTube"
              >
                <MusicIcon aria-hidden="true" className="size-4" />
              </SocialIconLink>
              {/* <DonateLink /> */}
              <ThemesGeneralSwitcher />
            </div>
            <div className="text-sm">
              © {new Date().getFullYear()}{" "}
              <Link
                href={"https://github.com/alifdwt"}
                rel="noreferrer noopener"
                target="_blank"
              >
                {env.NEXT_PUBLIC_APP_URL === "https://kutipanda.vercel.app" ? (
                  <>
                    {siteConfig.author.fullName}
                    <span className="sr-only">
                      {siteConfig.author.fullName}&apos;s GitHub Page
                    </span>
                  </>
                ) : (
                  <>{siteConfig.name}</>
                )}
              </Link>
              <br />
              <span className="font-normal">
                Hak cipta dilindungi undang-undang
              </span>
            </div>
          </div>
        </section>
        <section
          className={`
            grid flex-1 grid-cols-1 gap-10 text-center

            sm:grid-cols-4

            xxs:grid-cols-2
          `}
        >
          {siteConfig.footerNav.map((item) => (
            <div className="space-y-3" key={item.title}>
              <h4 className="text-base font-medium">{item.title}</h4>
              <ul className="space-y-2.5">
                {item.items.map((link) => (
                  <li key={link.title}>
                    <Link
                      className={`
                        text-sm text-muted-foreground transition-colors

                        hover:text-foreground
                      `}
                      href={link.href}
                      rel={link?.external ? "noreferrer" : undefined}
                      target={link?.external ? "_blank" : undefined}
                    >
                      {link.title}
                      <span className="sr-only">{link.title}</span>
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </section>
      </Shell>
    </footer>
  );
}
