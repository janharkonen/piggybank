import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card";
import { useAction } from "convex/react";
import { api } from "@/convex/_generated/api";
import { Button } from "@/components/ui/button";

export default function PopulateModal() {
  const populateSaldo = useAction(api.myFunctions.populateSaldo);
  return (
    <div className="flex flex-col gap-8  w-full max-w-5xl mx-auto">
    <Card>
      <CardHeader>
        <CardTitle>Lisää dataa</CardTitle>
      </CardHeader>
      <CardContent>
        <Button onClick={() => { populateSaldo() }}>Populate Saldo</Button>
      </CardContent>
      <CardFooter></CardFooter>
    </Card>
    </div>
  );
}