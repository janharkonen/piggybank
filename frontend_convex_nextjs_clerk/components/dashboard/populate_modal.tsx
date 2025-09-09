import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card";
import { useAction } from "convex/react";
import { api } from "@/convex/_generated/api";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function PopulateModal() {
  const populateSaldo = useAction(api.myFunctions.populateSaldo);
  return (

    <div className="flex flex-col gap-8  w-full max-w-5xl mx-auto">
    <Card>
      <CardHeader>
        <CardTitle>Lisää tiedosto</CardTitle>
      </CardHeader>
      <CardContent>
        <Input type="file" id="fileInput" className="mb-4" />
        <Button onClick={() => {
          const fileInput = document.getElementById('fileInput') as HTMLInputElement;
          const file = fileInput.files?.[0];
          if (file) {
            populateSaldo();
          } else {
            alert('Please select a file first');
          }
        }}>Populate Saldo</Button>
      </CardContent>
      <CardFooter></CardFooter>
    </Card>
    </div>
  );
}