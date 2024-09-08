"use server";

import { IQuote } from "@/interfaces/quote";
import { IResponse } from "@/interfaces/response";
import { BASE_URL } from "../BaseUrl";

export const getRandomQuotes = async (count: number) => {
  let data: IResponse<IQuote[]> | undefined = undefined;
  let isError = false;
  let error = "";

  try {
    const res = await fetch(`${BASE_URL}/quote/random/${count}`);
    data = await res.json();
  } catch (e) {
    isError = true;
    if (typeof e === "string") error = e;
    else if (e instanceof Error) error = e.message;
    else error = "Error";
  }
};
