import { BannerWithButton } from "@/components/Marketing/Elements/Banner/with-button";
import HomeHeroSection from "@/components/Marketing/HeroSection";

export default function Home() {
  return (
    <>
      <BannerWithButton
        linkHref="https://github.com/alifdwt/kutipanda"
        title="Lorem ipsum dolor sit amet!"
        detail="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc consequat."
        button="Get Started"
        dismiss="Dismiss"
      />
      <HomeHeroSection />
    </>
  );
}
