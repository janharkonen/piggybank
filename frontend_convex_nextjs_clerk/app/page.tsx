"use client";

import {
  Authenticated,
  Unauthenticated,
  useMutation,
  useQuery,
} from "convex/react";
import { api } from "../convex/_generated/api";
import Link from "next/link";
import { SignUpButton } from "@clerk/nextjs";
import { SignInButton } from "@clerk/nextjs";
import { UserButton } from "@clerk/nextjs";
import DarkModeToggle from "@/components/Header/DarkModeToggle";
import { Card, CardHeader, CardTitle, CardFooter, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import Nettoarvo from "@/components/dashboard/nettoarvo";

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
    border-b-2 
    bg-slate-100
    border-slate-300 
    dark:bg-slate-700
    dark:border-slate-600 
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
    /*
    <div className="
    w-full
    max-w-4xl
    mx-auto
    shadow-2xl
    border-2
    border-mint-200
    flex 
    flex-col 
    rounded-xl 
    "
    style={{ background: 'var(--foreground)' }}
    >
    <div className="logo-div h-full flex flex-col items-center justify-center"> 
    <img 
    src="https://picapi.janharkonen.fi/api/pics/224ad3d792204463bf57ae0eda3bebde.png?BG=92"
      alt="Piggybank Logo"
      style={{aspectRatio: 1 / 1, filter: "invert(0)"}}
      className="logo-img opacity-75 max-h-60 aspect-square"

      />
      <p className="font-bold text-3xl sm:text-5xl p-4">Piggybank</p>
      <Button className="bg-foreground text-background px-4 py-2 rounded-md">
      Sign in
        </Button>
        </SignInButton>
        <SignUpButton mode="modal">
        <button className="bg-foreground text-background px-4 py-2 rounded-md">
        Sign up
        </button>
        </SignUpButton>
        
        </div>
        </div>
        */
  );
}

function Content() {
  const { viewer, numbers } =
    useQuery(api.myFunctions.listNumbers, {
      count: 5,
    }) ?? {};
  const netValue = useQuery(api.myFunctions.netValue);
  const addNumber = useMutation(api.myFunctions.addNumber);

  if (viewer === undefined || numbers === undefined) {
    return (
      <div className="mx-auto">
        <p>loading... (consider a loading skeleton)</p>
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-8 w-full max-w-5xl mx-auto">
      <Nettoarvo viewer={viewer ?? "Anonymous"} netValue={netValue?.[0]?.value ?? 0} />
    </div>
  );
}

function ResourceCard({
  title,
  description,
  href,
}: {
  title: string;
  description: string;
  href: string;
}) {
  return (
    <div className="flex flex-col gap-2 bg-slate-200 dark:bg-slate-800 p-4 rounded-md h-28 overflow-auto">
      <a href={href} className="text-sm underline hover:no-underline">
        {title}
      </a>
      <p className="text-xs">{description}</p>
    </div>
  );
}
