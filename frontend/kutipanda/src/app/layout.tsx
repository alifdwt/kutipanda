import type { Metadata, Viewport } from "next";
import { Inter } from "next/font/google";
import "../styles/globals.css";
import { env } from "@/env";
import { siteConfig } from "@/app";
import { ThemeProvider } from "@/components/Providers/ThemeProvider";
import { SiteHeader } from "@/components/Navigation/SiteHeader";
import { SiteFooter } from "@/components/Navigation/SiteFooter";
import { cn } from "@/lib/utils";

const inter = Inter({ subsets: ["latin"] });

const baseMetadataURL =
  (env.NEXT_PUBLIC_APP_URL as string) || "http://localhost:3000";

export async function generateMetadata() {
  const metadata: Metadata = {
    alternates: {
      canonical: new URL(baseMetadataURL),
    },
    applicationName: siteConfig.name,
    authors: [{ name: siteConfig.author.fullName, url: siteConfig.author.url }],
    creator: siteConfig.author.handle,
    description: siteConfig.appNameDesc,
    icons: [{ rel: "icon", url: "/favicon.ico" }],
    keywords: siteConfig.keywords,
    other: {
      "darkreader-lock": "meta",
    },
    publisher: siteConfig.appPublisher,
    robots: {
      follow: true,
      index: true,
    },
    title: {
      default: siteConfig.name,
      template: `%s – ${siteConfig.appNameDesc}`,
    },
    twitter: {
      card: "summary_large_image",
      creator: siteConfig.author.handleAt,
      description: siteConfig.appNameDesc,
      images: [`${baseMetadataURL}/og.png`],
      site: siteConfig.author.handleAt,
      title: siteConfig.name,
    },
  };

  return metadata;
}

export const viewport: Viewport = {
  colorScheme: "dark light",
  themeColor: [
    { color: "white", media: "(prefers-color-scheme: light)" },
    { color: "black", media: "(prefers-color-scheme: dark)" },
  ],
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={cn(
          "min-h-screen bg-background font-sans antialiased",
          inter.className
        )}
      >
        <ThemeProvider>
          <SiteHeader />
          {children}
          <SiteFooter />
        </ThemeProvider>
      </body>
    </html>
  );
}
