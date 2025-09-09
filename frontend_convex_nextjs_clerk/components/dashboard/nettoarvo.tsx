import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card";


export default function Nettoarvo({ viewer, netValue }: { viewer: string, netValue: number }) {
  return (

    <div className="flex flex-col gap-8  w-full max-w-5xl mx-auto">
    <Card>
      <CardHeader>
        <CardTitle>Welcome {viewer ?? "Anonymous"}!</CardTitle>
      </CardHeader>
      <CardContent>
        <p>Nettoarvo: {netValue}</p>
      </CardContent>
      <CardFooter></CardFooter>
    </Card>
    </div>
  );
}