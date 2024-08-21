import { FooterItem } from "../addons/types/kutipanda/nav";
import { MainMenuItem } from "../addons/types/kutipanda/with";
import metadata from "./constants/metadata";

const socialLinks = {
  instagram: "https://www.instagram.com/kutipanda",
  twitter: "https://twitter.com/kutipanda",
  github: "https://github.com/alifdwt/kutipanda",
  githubAccount: "https://github.com/alifdwt",
  discord: "https://discord.gg/kutipanda",
  telegram: "https://t.me/kutipanda",
};

export const siteConfig = {
  name: metadata.name,
  appNameDesc: metadata.appNameDesc,
  appPublisher: metadata.appPublisher,
  appVersion: metadata.appVersion,
  author: {
    email: metadata.author.email,
    fullName: metadata.author.fullName,
    handle: metadata.author.handle,
    handleAt: metadata.author.handleAt,
    url: metadata.author.url,
  },
  footerNav: [
    {
      items: [
        {
          external: false,
          href: "/contact",
          title: "Kontak",
        },
        {
          external: false,
          href: "/privacy",
          title: "Privacy",
        },
        {
          external: false,
          href: "/terms",
          title: "Ketentuan",
        },
        {
          external: false,
          href: "/about",
          title: "Tentang Kami",
        },
      ],
      title: "Bantuan",
    },
    {
      items: [
        {
          external: true,
          href: socialLinks.instagram,
          title: "Instagram",
        },
        {
          external: true,
          href: socialLinks.twitter,
          title: "Twitter",
        },
        {
          external: true,
          href: socialLinks.github,
          title: "GitHub",
        },
        {
          external: true,
          href: socialLinks.discord,
          title: "Discord",
        },
        {
          external: true,
          href: socialLinks.telegram,
          title: "Telegram",
        },
      ],
      title: "Media Sosial",
    },
    {
      items: [
        {
          external: true,
          href: "https://github.com/alifdwt/kutipanda",
          title: "@kutipanda",
        },
        {
          external: true,
          href: socialLinks.githubAccount,
          title: "@alifdwt",
        },
        {
          external: true,
          href: socialLinks.github,
          title: "Alif Dewantara",
        },
      ],
      title: "Github",
    },
    {
      items: [
        {
          external: true,
          href: "https://github.com/sponsors/alifdwt",
          title: "GitHub Sponsors",
        },
        {
          external: true,
          href: "https://buymeacoffee.com/alifdwt",
          title: "Buy Me a Coffee",
        },
        {
          external: true,
          href: "https://patreon.com/alifdwt",
          title: "Patreon",
        },
        {
          external: true,
          href: "https://paypal.me/alifdwt",
          title: "PayPal",
        },
      ],
      title: "Support",
    },
  ] satisfies FooterItem[],
  images: {
    alt: "Show the cover image of Kutipanda",
    url: "/og.png",
  },
  keywords: ["lyric", "hindi", "arab", "indonesia"] as string[],
  links: socialLinks,
  mainNav: [
    {
      href: "/",
      items: [
        {
          title: "Film",
          href: "/movies",
          items: [],
          description: "Ulasan, rekomendasi, dan katalog film",
        },
        {
          title: "Lagu",
          href: "/songs",
          items: [],
          description: "Lagu terbaik yang kamu suka",
        },
        {
          title: "Kutipan",
          href: "/quotes",
          items: [],
          description: "Kata-kata terindah",
        },
      ],
      title: "Katalog",
    },
    {
      href: "/blog",
      items: [
        {
          title: "Blog Film",
          href: "/blog/categories/movies",
          items: [],
          description: "Tulisan terkini tentang film",
        },
        {
          title: "Blog Lagu",
          href: "/blog/categories/songs",
          items: [],
          description: "Tulisan terkini tentang lagu",
        },
        {
          title: "Blog Kutipan",
          href: "/blog/categories/quotes",
          items: [],
          description: "Tulisan terkini tentang kutipan",
        },
        {
          title: "Blog Tentang Kami",
          href: "/blog/categories/about",
          items: [],
          description: "Tulisan terkini tentang tentang kami",
        },
      ],
      title: "Blog",
    },
  ] satisfies MainMenuItem[],
  socialLinks: {
    github: socialLinks.github,
    instagram: socialLinks.instagram,
    twitter: socialLinks.twitter,
  },
  themeToggleEnabled: true,
} as const;
