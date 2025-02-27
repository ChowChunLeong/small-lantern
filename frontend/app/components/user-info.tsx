import { auth, signOut } from "../auth";
import ProfileImage from "./profile-image";

export default async function UserProfile() {
  const session = await auth();

  if (!session?.user) return null;
  const srcImage = session.user.image ?? "";

  return (
    <div>
      <form
        action={async () => {
          "use server";
          await signOut({ redirectTo: "/login" });
        }}
      >
        <button type="submit">Sign Out</button>
      </form>
      <ProfileImage imageUrl={srcImage} alt={"User Profile"} />
    </div>
  );
}
