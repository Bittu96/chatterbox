import { auth, signOut } from "@/app/api/auth/[...nextauth]/route";
import { redirect } from "next/navigation";
import NavBar from "@/app/components/navbar/nav-bar";

export default async function ProtectedPage() {
  console.log("checking session");
  let session = await auth();
  console.log("session", session, session?.user?.email == undefined);

  if (session?.user?.email == undefined) {
    redirect("/login");
  }
  
  return (
    <div>
      <NavBar profile={session?.user} />

      <div className="flex h-screen w-screen items-center justify-center">
        <div className="z-10 w-full max-w-md  rounded-2xl border border-gray-100 shadow-xl bg-white">
          {/* <div className="flex flex-col items-center justify-center space-y-3 border-b border-gray-200 bg-white px-4 py-6 pt-8 text-center sm:px-16">
            <h3 className="text-xl font-semibold">Chatterbox</h3>
          </div> */}

          <div className="flex flex-col space-y-4  px-4 py-8 sm:px-16">
            You are logged in as {session?.user?.email}
            <SignOut />
          </div>

          <div className="items-center justify-center  border-b border-gray-200 bg-white py-6 pt-8 text-center"></div>
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
      <button className="button" type="submit">Sign out</button>
    </form>
  );
}
