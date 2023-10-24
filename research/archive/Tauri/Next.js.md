# Next.js
Per https://nextjs.org/learn/basics/create-nextjs-app/setup

```
npx create-next-app nextjs-blog --use-npm --example "https://github.com/vercel/next-learn/tree/master/basics/learn-starter"
```

Then:
```
cd nextjs-blog
npm run dev
```

Then in another window:
```
npm install --save-dev @tauri-apps/cli
npm tauri init # did not work
yarn tauri init
yarn tauri dev
```

Worked ok.
