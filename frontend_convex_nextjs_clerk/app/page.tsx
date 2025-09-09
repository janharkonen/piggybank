"use client";

import {
  Authenticated,
  Unauthenticated,
} from "convex/react";
import { SignUpButton } from "@clerk/nextjs";
import { SignInButton } from "@clerk/nextjs";
import { UserButton } from "@clerk/nextjs";
import DarkModeToggle from "@/components/Header/DarkModeToggle";
import { Card, CardHeader, CardTitle, CardFooter } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import Nettoarvo from "@/components/dashboard/nettoarvo";
import PopulateModal from "@/components/dashboard/populate_modal";

export default function Home() {
  return (
    <>
      <main className="flex flex-col h-screen flex-grow">
        <Authenticated>
          <Header />
          <Content />
        </Authenticated>
        <Unauthenticated>
        <div className="
          w-full
          h-screen
          flex
          justify-center
          items-center
        ">
          <div className="absolute top-0 right-0">
            <DarkModeToggle />
          </div>
          <SignInForm />
        </div>
        </Unauthenticated>
      </main>
    </>
  );
}

function Header() {
  return (
    <header 
    className="
    sticky 
    top-0 
    z-10 
    p-2 
    bg-[var(--header-background)] 
    border-b-[var(--header-border)]
    border-b-2 
    flex 
    flex-row 
    gap-2
    items-center"
    >
      <span className="text-lg font-bold pl-2 flex-grow">
        Piggybank
      </span>
      <div className="flex-none" ><DarkModeToggle/></div>
      <div className="flex flex-col items-center" ><UserButton/></div>
    </header>
  );
}

function SignInForm() {
  return (
    <Card className="w-full max-w-xl">
      <div className="relative top-0 right-4 self-end justify-self-end">
        <SignUpButton mode="modal">
          <Button variant="link" className="hover:cursor-pointer">Sign Up</Button>
        </SignUpButton>
      </div>
      <img 
      src="https://picapi.janharkonen.fi/api/pics/224ad3d792204463bf57ae0eda3bebde.png?BG=92"
      alt="Piggybank Logo"
      style={{
        aspectRatio: 1 / 1,
      }}
      className="logo-img opacity-75 max-h-60 max-w-60 w-full aspect-square self-center filter dark:invert"
      />
    <CardHeader>
      <CardTitle className="text-3xl font-bold text-center">Piggybank</CardTitle>
      {/*
    */}
    </CardHeader>
    <CardFooter className="flex-col gap-2">
      <SignInButton mode="modal">
        <Button type="submit" className="w-full hover:cursor-pointer">Login</Button>
      </SignInButton>
    </CardFooter>
  </Card>
  );
}

function Content() {

  return (
    <div className="flex flex-col gap-8 w-full max-w-5xl mx-auto">
      <Nettoarvo />
      <PopulateModal />
    </div>
  );
}