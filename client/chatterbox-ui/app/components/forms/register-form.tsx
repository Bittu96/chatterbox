import { redirect } from "next/navigation";
import { SocialButton } from "@/app/components/buttons/social-button";

export function RegisterForm({
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
        className="flex flex-col space-y-4 px-4 py-8 sm:px-16"
      >
        <div>
          <input
            id="username"
            name="username"
            type="text"
            placeholder="username"
            required
            className="mt-1 block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-black focus:outline-none focus:ring-black sm:text-sm"
          />
        </div>
        <div>
          <input
            id="email"
            name="email"
            type="email"
            placeholder="user@box.com"
            autoComplete="email"
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
        <div>
          <input
            id="confirm-password"
            name="confirm-password"
            type="password"
            placeholder="confirm password"
            required
            className="mt-1 block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 placeholder-gray-400 shadow-sm focus:border-black focus:outline-none focus:ring-black sm:text-sm"
          />
        </div>
        {children}
      </form>

      <div className="flex flex-col text-center sm:px-16 text-sm text-gray-400">
        or you can sign in with
      </div>

      <form
        action={async (formData: FormData) => {
          "use server";
          redirect("/api/auth/signin/google");
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
