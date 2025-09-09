import { v } from "convex/values";
import { query, mutation, action } from "./_generated/server";
import { api } from "./_generated/api";

// Write your Convex functions in any file inside this directory (`convex`).
// See https://docs.convex.dev/functions for more.

export const netValue = query({
  handler: async (ctx) => {
    const netValue = await ctx.db
      .query("nettoarvo")
      .take(1);
    return netValue;
  },
});

// You can read data from the database via a query:
export const listNumbers = query({
  // Validators for arguments.
  args: {
    count: v.number(),
  },

  // Query implementation.
  handler: async (ctx, args) => {
    //// Read the database as many times as you need here.
    //// See https://docs.convex.dev/database/reading-data.
    console.log("--------------------------------------")
    const identity = await ctx.auth.getUserIdentity()
    console.log(identity)
    const numbers = await ctx.db
      .query("numbers")
      // Ordered by _creationTime, return most recent
      .order("desc")
      .take(args.count);
    return {
      viewer: (await ctx.auth.getUserIdentity())?.email ?? null,
      numbers: numbers.reverse().map((number) => number.value),
    };
  },
});

// You can write data to the database via a mutation:
export const addNumber = mutation({
  // Validators for arguments.
  args: {
    value: v.number(),
  },

  // Mutation implementation.
  handler: async (ctx, args) => {
    //// Insert or modify documents in the database here.
    //// Mutations can also read from the database like queries.
    //// See https://docs.convex.dev/database/writing-data.

    const id = await ctx.db.insert("numbers", { value: args.value });

    console.log("Added new document with id:", id);
    // Optionally, return a value from your mutation.
    // return id;
  },
});

// You can fetch data from and send data to third-party APIs via an action:
export const myAction = action({
  // Validators for arguments.
  args: {
    first: v.number(),
    second: v.string(),
  },

  // Action implementation.
  handler: async (ctx, args) => {
    //// Use the browser-like `fetch` API to send HTTP requests.
    //// See https://docs.convex.dev/functions/actions#calling-third-party-apis-and-using-npm-packages.
    // const response = await ctx.fetch("https://api.thirdpartyservice.com");
    // const data = await response.json();

    //// Query data by running Convex queries.
    const data = await ctx.runQuery(api.myFunctions.listNumbers, {
      count: 10,
    });
    console.log(data);

    //// Write data by running Convex mutations.
    await ctx.runMutation(api.myFunctions.addNumber, {
      value: args.first,
    });
  },
});

type KuukausiSaldoRow ={
  kategoria: string;
  kuka: string;
  palvelu: string;
  mita: string;
  aikaleima: string;
  erittely: string;
  saldo: number;
}

async function getFromGoogleDocs(userid: string) {
  const gvizUrl = process.env.CONVEX_POPULATE_URL1 as string;
  console.log("gvizUrl", gvizUrl);
  const response = await fetch(gvizUrl);
  const txt = await response.text();
  const json = JSON.parse(txt.match(/setResponse\((.*)\)/)?.[1] ?? '');
  const table = json.table;
  //console.log("table", table);
  const rows = table.rows.map((row: any) => row.c.map((cell: any) => cell?.v ?? null));
  const columns = table.cols.map((col: any) => col.label);
 
  var kuukausiSaldoRow: KuukausiSaldoRow = {
    kategoria: "",
    kuka: "",
    palvelu: "",
    mita: "",
    aikaleima: "",
    erittely: "",
    saldo: 0,
  };


  var kuukausiSaldoRows: KuukausiSaldoRow[] = [];
  kuukausiSaldoRow.kategoria = "Sijoitukset";
  kuukausiSaldoRow.kuka = userid;
  kuukausiSaldoRow.erittely = "";
  rows.slice(0, 1).forEach((row: any, rowIndex: number) => {
    kuukausiSaldoRow.palvelu = row[0];
    kuukausiSaldoRow.mita = row[1];
    console.log("tassa: ", columns)
    columns.forEach((value: any, colIndex: number) => {
      if (colIndex > 2) {
        //console.log("column index: ", colIndex)
        //console.log("column value: ", value)
        //console.log('value type:', typeof value)
        kuukausiSaldoRow.aikaleima = (value as string).trim();
        kuukausiSaldoRow.saldo = parseFloat(row[colIndex] as string);
        kuukausiSaldoRows = [...kuukausiSaldoRows, {...kuukausiSaldoRow}];
      }
    });
  });

  return kuukausiSaldoRows;
};

export const insertKuukausiSaldo = mutation({
  args: {
    row: v.object({
      kategoria: v.string(),
      kuka: v.string(),
      palvelu: v.string(),
      mita: v.string(),
      aikaleima: v.string(),
      erittely: v.string(),
      saldo: v.number(),
    }),
  },

  handler: async (ctx, args) => {
    await ctx.db.insert("kuukausi_saldo", args.row);
  },
});

export const insertKuukausiSaldoArray = mutation({
  args: {
    rows: v.array(v.object({
      kategoria: v.string(),
      kuka: v.string(),
      palvelu: v.string(),
      mita: v.string(),
      aikaleima: v.string(),
      erittely: v.string(),
      saldo: v.number(),
    })),
  },

  handler: async (ctx, args) => {
    args.rows.forEach(async (row: KuukausiSaldoRow) => {
      await ctx.runMutation(api.myFunctions.insertKuukausiSaldo, {
        row: row,
      });
    });
  },
});

export const populateSaldo = action({
  handler: async (ctx) => {
    const user = await ctx.auth.getUserIdentity();
    console.log("user", user);
    const userid = (user?.tokenIdentifier ?? "").split("|")[1];
    const data = await getFromGoogleDocs(userid);
    await ctx.runMutation(api.myFunctions.insertKuukausiSaldoArray, {
      rows: data,
    });
    
    
  },
})