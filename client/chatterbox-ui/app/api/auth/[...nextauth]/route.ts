import GoogleProvider from "next-auth/providers/google";
import NextAuth from "next-auth";

export const authOptions = {
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET,
    }),
  ],
  // callbacks: {
  //   async signIn({ account, profile }:any) {
  //     console.log("account:",account)
  //     console.log("profile:",profile)

  //     return redirect("/")
  //     if (account.provider === "google") {
  //       return profile.email_verified && profile.email.endsWith("@gmail.com")
  //     }
  //     return redirect("/")
  //     return true // Do different verification for other providers that don't have `email_verified`
  //   },
  // }
};

export const {
  handlers: { GET, POST },
  auth,
  signIn,
  signOut,
} = NextAuth(authOptions);
