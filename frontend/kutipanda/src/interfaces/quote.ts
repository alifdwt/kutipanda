import { ILanguage } from "./language";
import { IMovie } from "./movie";
import { IUser } from "./user";

export interface IQuote {
  id: number;
  quote_text: string;
  quote_transliteration: string;
  language_id: number;
  language: ILanguage;
  movie_id: number;
  movie: IMovie;
  user_id: number;
  user: IUser;
  created_at: string;
  updated_at: string;
}
