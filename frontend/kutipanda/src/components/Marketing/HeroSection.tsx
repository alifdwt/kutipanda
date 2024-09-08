import HeroQuotes from "./Elements/HeroElements/HeroQuotes";

export default function HomeHeroSection() {
  return (
    <section className="grid grid-rows-4 grid-cols-5 grid-flow-col gap-4 *:rounded-xl">
      <div className="min-h-96 bg-blue-200 row-span-4 col-span-3">
        <HeroQuotes />
      </div>
      <div className="bg-green-200 col-span-2 row-span-3"></div>
      <div className="bg-green-700 col-span-2"></div>
    </section>
  );
}
