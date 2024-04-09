import Link from "next/link";
import { LoginForm } from "@/app/components/forms/login-form";
import { SubmitButton } from "@/app/components/buttons/submit-button";
import { AuthError } from "next-auth";
import { redirect } from "next/navigation";

export default function Login() {
  return (
    <div className="flex flex-col h-screen w-screen items-center justify-center">
      <h3 className="title text-5xl text-white mb-2">Chatterbox</h3>

      <div className="z-10 w-full max-w-md overflow-hidden rounded-2xl border border-gray-100 shadow-xl bg-white">
        <div className="flex flex-col items-center justify-center space-y-3 bg-white px-4 py-4 pt-8 text-center sm:px-16"></div>

        <LoginForm
          action={async (formData: FormData) => {
            "use server";
            logInAction(formData);
          }}
        >
          <SubmitButton>Sign in</SubmitButton>
          <p className="text-center text-sm text-gray-600">
            {"Don't have an account? "}
            <Link href="/register" className="font-semibold text-gray-800">
              Sign up
            </Link>
            {" for free."}
          </p>
        </LoginForm>

        <div className="items-center justify-center  border-b border-gray-200 bg-white py-6 pt-8 text-center"></div>
      </div>
    </div>
  );
}

export async function logInAction(formData: FormData) {
  redirect("/");

  // try {
  //   await signIn("credentials", {
  //     // redirect:false,
  //     redirectTo: "http://localhost:3000/protected",
  //     name: formData.get("username") as string,
  //     password: formData.get("password") as string,
  //   });
  // } catch (error) {
  //   console.log("found error");

  //   if (error instanceof AuthError) {
  //     console.log("error.type",error.type);

  //     switch (error.type) {
  //       case "CredentialsSignin":
  //         console.log( "Invalid credentials");
  //         return { error: "Invalid credentials" };
  //       default:
  //         console.log("Something went wrong");
  //         // redirect("http://localhost:3000/protected")
  //         return { error: "Something went wrong" };
  //     }
  //   }
  //    else {
  //     redirect("http://localhost:3000/protected")
  //   }
  //   // else {
  //   //   console.log(error);
  //   //   throw error;
  //   // }
  // }
}
