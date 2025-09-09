import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card";
import {
  useQuery,
} from "convex/react";
import { api } from "@/convex/_generated/api";

export default function Nettoarvo() {
  const { viewer, numbers } =
    useQuery(api.myFunctions.listNumbers, {
      count: 5,
    }) ?? {};
  const netValue = useQuery(api.myFunctions.netValue);
  
  if (viewer === undefined || numbers === undefined) {
    return (
      <div className="mx-auto">
        <p>loading... (consider a loading skeleton)</p>
      </div>
    );
  }
  return (

    <div className="flex flex-col gap-8  w-full max-w-5xl mx-auto">
    <Card>
      <CardHeader>
        <CardTitle>Welcome {viewer ?? "Anonymous"}!</CardTitle>
      </CardHeader>
      <CardContent>
        <p>Nettoarvo: {netValue?.[0]?.value ?? 0}</p>
      </CardContent>
      <CardFooter></CardFooter>
    </Card>
    </div>
  );
}