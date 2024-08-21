import { SidebarNavItem } from "../../addons/types/kutipanda/with";

type DashboardConfig = {
  sidebarNav: SidebarNavItem[];
};

export const dashboardConfig: DashboardConfig = {
  sidebarNav: [
    {
      href: "/dashboard",
      icon: "dollarSign",
      items: [],
      title: "Dashboard",
    },
  ],
};
