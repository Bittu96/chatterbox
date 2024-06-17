"use server";

import Cookies from "js-cookie";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";

export default async function SetCookie() {
  console.log("asegsg");
  let c = Cookies.set("session-username", "some-value-by-client-comp", {
    expires: 200000,
  });
  console.log(c)
  cookies().set('session-username', 'kittu')
  redirect("/chat")
}
