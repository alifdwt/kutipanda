import { ICountry } from "./country";
import { ISong } from "./song";

export interface IMovie {
  id: number;
  title: string;
  slug: string;
  description: string;
  release_date: string;
  poster_image_url: string;
  country_id: number;
  country: ICountry;
  // actors: IActor[];
  songs: ISong[];
  created_at: string;
  updated_at: string;
}
