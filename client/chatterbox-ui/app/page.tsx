import { auth, signOut } from "@/app/api/auth/[...nextauth]/route";
import { redirect } from "next/navigation";
import Home from "@/app/components/home/page";
import { SubmitButton } from "@/app/components/buttons/submit-button";
import SignOutDialog from "./components/dialogs/page";

export default async function Page() {
  console.log("checking session");
  let session = await auth();
  console.log("session", session, session?.user?.email == undefined);

  if (session?.user?.email == undefined) {
    redirect("/login");
  }

  const profile = {
    name: session.user.name,
    email: session.user.email,
    role: "Co-Founder / CEO",
    image: session.user.image,
    lastSeen: "3h ago",
    lastSeenDateTime: "2023-01-23T13:23Z",
  };

  return (
    <div className="flex  flex-col h-screen overflow-auto items-center justify-center">
      <h3 className="title text-5xl text-white mb-2">Chatterbox</h3>
      <div className="z-10 w-full max-w-md overflow-auto rounded-2xl border border-gray-100 shadow-xl bg-white py-6">
        <div className="flex flex-col items-center justify-center bg-white text-center mb-5 sm:px-16">
          <p className="text-xl font-semibold leading-6 text-gray-900 title">
            Welcome, {profile.name}!
          </p>
        </div>

        <div className="flex menu-bar" role="group">
          <button
            type="button"
            className="inline-flex items-center text-sm font-medium hover:text-white"
          >
            <svg
              className="h-3 me-2"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path d="M10 0a10 10 0 1 0 10 10A10.011 10.011 0 0 0 10 0Zm0 5a3 3 0 1 1 0 6 3 3 0 0 1 0-6Zm0 13a8.949 8.949 0 0 1-4.951-1.488A3.987 3.987 0 0 1 9 13h2a3.987 3.987 0 0 1 3.951 3.512A8.949 8.949 0 0 1 10 18Z" />
            </svg>
            Profile
          </button>
          <button
            type="button"
            className="inline-flex items-center text-sm font-medium hover:text-white"
          >
            <svg
              className="h-3 me-2"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 20 20"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 12.25V1m0 11.25a2.25 2.25 0 0 0 0 4.5m0-4.5a2.25 2.25 0 0 1 0 4.5M4 19v-2.25m6-13.5V1m0 2.25a2.25 2.25 0 0 0 0 4.5m0-4.5a2.25 2.25 0 0 1 0 4.5M10 19V7.75m6 4.5V1m0 11.25a2.25 2.25 0 1 0 0 4.5 2.25 2.25 0 0 0 0-4.5ZM16 19v-2"
              />
            </svg>
            Home
          </button>
          <button
            type="button"
            className="inline-flex items-center text-sm font-medium hover:text-white"
          >
            <svg
              className=" h-3 me-2"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path d="M14.707 7.793a1 1 0 0 0-1.414 0L11 10.086V1.5a1 1 0 0 0-2 0v8.586L6.707 7.793a1 1 0 1 0-1.414 1.414l4 4a1 1 0 0 0 1.416 0l4-4a1 1 0 0 0-.002-1.414Z" />
              <path d="M18 12h-2.55l-2.975 2.975a3.5 3.5 0 0 1-4.95 0L4.55 12H2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2Zm-3 5a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z" />
            </svg>
            Chat
          </button>
        </div>

        <hr />

        <Home />

        <hr />

        <div className="items-center justify-center pt-6 px-16 text-center">
          {/* <SignOut /> */}
          <SignOutDialog params={signOutAction} />
        </div>
      </div>
    </div>
  );
}

function SignOut() {
  return (
    <form
      action={async () => {
        "use server";
        await signOut();
      }}
    >
      <SubmitButton>Sign out</SubmitButton>
    </form>
  );
}

const signOutAction = async function () {
  "use server";
  console.log("sign out called");
  await signOut({ redirect: true, redirectTo: "/" });
};
