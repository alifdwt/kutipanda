import { ICountry } from "./country";
import { ILanguage } from "./language";
import { IMovie } from "./movie";
import { IUser } from "./user";

export interface ISong {
  id: number;
  title: string;
  slug: string;
  description: string;
  release_date: string;
  album_image_url: string;
  user_id: number;
  user?: IUser;
  country_id: number;
  country?: ICountry;
  language_id: number;
  language?: ILanguage;
  // artists: IArtist[];
  movies?: IMovie[];
  created_at: string;
  updated_at: string;
}
