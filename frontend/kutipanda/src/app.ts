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
          title: "Privasi",
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
      title: "Unggulan",
    },
    {
      title: "Film",
      href: "/movies",
      items: [
        {
          title: "Rekomendasi Film",
          href: "/movies/recommendations",
          items: [],
          description: "Rekomendasi film dari kami",
        },
        {
          title: "Genre",
          href: "/movies/genres",
          items: [],
          description: "Genre film yang tersedia",
        },
        {
          title: "Tahun",
          href: "/movies/years",
          items: [],
          description: "Tahun film yang tersedia",
        },
        {
          title: "Negara",
          href: "/movies/countries",
          items: [],
          description: "Negara film yang tersedia",
        },
      ],
      featuredItem: {
        href: "/movies/kalank",
        title: "Andhadhun",
        description: "Ketika seorang pianis buta melihat terlalu banyak hal.",
        items: [],
      },
    },
    {
      title: "Lagu",
      href: "/songs",
      items: [
        {
          title: "Rekomendasi Lagu",
          href: "/songs/recommendations",
          items: [],
          description: "Rekomendasi lagu dari kami",
        },
        {
          title: "Genre",
          href: "/songs/genres",
          items: [],
          description: "Genre lagu yang tersedia",
        },
        {
          title: "Tahun",
          href: "/songs/years",
          items: [],
          description: "Tahun lagu yang tersedia",
        },
        {
          title: "Negara",
          href: "/songs/countries",
          items: [],
          description: "Negara lagu yang tersedia",
        },
      ],
      featuredItem: {
        href: "/songs/naina-da-kya-kasoor-andhadhun",
        title: "Naina Da Kya Kasoor (Andhadhun)",
        items: [],
        description: "Apa salahnya si mata ini?",
      },
    },
    {
      title: "Blog",
      href: "/blog",
      items: [
        {
          title: "Blog",
          href: "/blog",
          items: [],
          description: "Tulisan pribadi",
        },
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
      ],
      featuredItem: {
        href: "/blog/industri-kreatif-indonesia-identitas-jawabannya",
        title:
          "Apa yang kurang dari industri kreatif Indonesia? Identitas Jawabannya",
        description:
          "Ketika identitas sebuah bangsa dapat menjadi nilai eksotis dalam industri kreatif",
        items: [],
      },
    },
    {
      title: "Dukung",
      href: "/support",
    },
  ] satisfies MainMenuItem[],
  socialLinks: {
    github: socialLinks.github,
    instagram: socialLinks.instagram,
    twitter: socialLinks.twitter,
  },
  themeToggleEnabled: true,
} as const;
