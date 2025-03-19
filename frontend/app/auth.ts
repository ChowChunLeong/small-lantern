import NextAuth from "next-auth";
import Google from "next-auth/providers/google";
import { AUTH } from "./constants/auth";
import { BACKEND } from "./constants/backend";
import { OAuthRequest } from "./form/auth";

declare module "next-auth" {
  interface Session {
    error?: "RefreshTokenError";
  }
}

declare module "next-auth" {
  interface JWT {
    access_token: string;
    expires_at: number;
    refresh_token?: string;
    error?: "RefreshTokenError";
  }
}

export const { handlers, signIn, signOut, auth } = NextAuth({
  providers: [
    Google({
      // Google requires "offline" access_type to provide a `refresh_token`
      authorization: { params: { access_type: "offline", prompt: "consent" } },
    }),
  ],
  callbacks: {
    async jwt({ token, account }) {
      console.log("token", token);
      console.log("account", account);
      if (account) {
        try {
          const requestData: OAuthRequest = {
            provider: "GOOGLE",
            email: token.email as string,
            provider_account_id: account.providerAccountId,
            name: token.name as string,
            image: token.picture as string,
          };

          const response = await fetch(BACKEND.URL + "/api/auth/oauth", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(requestData),
          });

          const result = await response.json();
          console.log("API Response:", result);
        } catch (error) {
          console.error("Error calling API:", error);
        }
        // First-time login, save the `access_token`, its expiry and the `refresh_token`
        return {
          ...token,
          access_token: account.access_token,
          expires_at: account.expires_at,
          refresh_token: account.refresh_token,
        };
      } else if (
        //type guard
        typeof token.expires_at === "number" &&
        Date.now() < token.expires_at * 1000
      ) {
        // Subsequent logins, but the `access_token` is still valid
        return token;
      } else {
        // Subsequent logins, but the `access_token` has expired, try to refresh it
        if (!token.refresh_token) throw new TypeError("Missng refresh_token");

        try {
          if (typeof token.refresh_token !== "string") {
            throw new TypeError("Invalid refresh_token");
          }

          // The `token_endpoint` can be found in the provider's documentation. Or if they support OIDC,
          // at their `/.well-known/openid-configuration` endpoint.
          // i.e. https://accounts.google.com/.well-known/openid-configuration
          const response = await fetch("https://oauth2.googleapis.com/token", {
            method: "POST",
            body: new URLSearchParams({
              client_id: AUTH.GOOGLE_ID,
              client_secret: AUTH.GOOGLE_SECRET,
              grant_type: "refresh_token",
              refresh_token: token.refresh_token!,
            }),
          });

          const tokensOrError = await response.json();

          if (!response.ok) throw tokensOrError;

          const newTokens = tokensOrError as {
            access_token: string;
            expires_in: number;
            refresh_token?: string;
          };

          return {
            ...token,
            access_token: newTokens.access_token,
            expires_at: Math.floor(Date.now() / 1000 + newTokens.expires_in),
            // Some providers only issue refresh tokens once, so preserve if we did not get a new one
            refresh_token: newTokens.refresh_token
              ? newTokens.refresh_token
              : token.refresh_token,
          };
        } catch (error) {
          console.error("Error refreshing access_token", error);
          // If we fail to refresh the token, return an error so we can handle it on the page
          token.error = "RefreshTokenError";
          return token;
        }
      }
    },
    async session({ session, token }) {
      return {
        ...session,
        error: token.error as "RefreshTokenError" | undefined,
      };
    },
  },
});
