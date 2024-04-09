import { redirect } from "next/navigation";
import { SocialButton } from "@/app/components/buttons/social-button";
import { signIn } from "@/app/api/auth/[...nextauth]/route";

export function LoginForm({
  action,
  children,
}: {
  action: any;
  children: React.ReactNode;
}) {
  return (
    <div>
      <form
        action={action}
        className="flex flex-col space-y-4  px-4 py-8 sm:px-16"
      >
        <div>
          <input
            id="username"
            name="username"
            type="text"
            placeholder="username"
            autoComplete="username"
            required
            className="mt-1 block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-black focus:outline-none focus:ring-black sm:text-sm"
          />
        </div>
        <div>
          <input
            id="password"
            name="password"
            type="password"
            placeholder="password"
            required
            className="mt-1 block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-black focus:outline-none focus:ring-black sm:text-sm"
          />
        </div>
        {children}
      </form>

      <div className="inline-flex items-center justify-center w-full">
        <hr className="w-64 h-px my-8 bg-gray-300 border-0 dark:bg-gray-700" />
        <span className="absolute px-3 text-sm text-gray-400 -translate-x-1/2 bg-white left-1/2">
          or you can sign in with
        </span>
      </div>

      <form
        action={async (formData: FormData) => {
          "use server";
          await signIn("google", {redirectTo:"/"});
        }}
        className="sm:px-16"
      >
        <SocialButton>
          <img
            className="google"
            loading="lazy"
            height="24"
            width="24"
            id="provider-logo"
            src="https://authjs.dev/img/providers/google.svg"
          />
        </SocialButton>
      </form>
    </div>
  );
}
