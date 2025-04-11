# 🌐 VaporNet

> _"Best viewed on a haunted CRT at 3AM with MIDI autoplay on and your dreams slightly scrambled."_

**VaporNet** is a surreal HTTP client proxy that generates a procedurally hallucinated internet, one request at a time. Every page is a love letter to the Weird Wide Web of the late '90s and early 2000s—reborn through the mind of an LLM and dripping with glittery GIFs, cursed fonts, and synthetic nostalgia.

## 🧪 Usage (Coming Soon™)

Eventually you’ll be able to:

```bash
docker run -p 8080 --env OPENAI_API_KEY="<your API key>" ghrc.io/ivy/vapornet
```

Then browse the hallucinet by setting your browser to proxy to `http://localhost:8080`.
Try any URL you want. Nothing is off-limits.

## 🌀 What Is It?

VaporNet turns any web request into a portal to an alternate timeline where:
- GeoCities never died.
- Everyone has a fan page.
- The moon landing was faked by an interdimensional record label.
- The internet is a dreaming organism held together by webrings and good intentions.

It works like this:
1. A browser makes a request to any URL (e.g. `http://vampireapartment.net/blood_ring/fanfic23.html`)
2. VaporNet interprets the URL, headers, and context as a **prompt**.
3. A large language model (LLM) hallucinates what kind of site that would be in this alternate reality.
4. VaporNet generates the HTML on the fly, complete with:
   - Imaginary assets (GIFs, MIDI, PDFs, broken webrings)
   - Nostalgic CSS and DHTML weirdness
   - Fictional webmasters and lost net lore
   - Pop-up greetings, mouse trails, `<marquee>` chaos, and other cursed artifacts

## 💾 Why?

Because the modern internet is too sterile.
Because we miss the chaos.
Because memory is a file format.
Because you deserve to browse a fake Sailor Moon fan club hosted on a server powered by moonlight and dial-up screams.

## ✨ Features

- 🌈 Full HTML hallucinations for any domain or path
- 📼 Fake assets dynamically generated when requested
- 💌 Interlinked surreal websites and user personas
- 🔮 DHTML effects and cursed JS hacks from a forgotten era
- 🧠 Modular LLM prompt pipeline (concept > HTML > assets)

## 🚧 How It Works

VaporNet is composed of:

- **An HTTP proxy server** that intercepts requests and parses the URL + metadata into prompts.
- **Prompt Engine**:
  - First prompt: generates a high-level *concept* for the page.
  - Second prompt: turns the concept into fully styled HTML.
  - Additional prompts for assets (images, audio, PDFs, etc.).
- **Asset Generator**: returns fake media when requested via URLs like `http://vapornet.net/assets/images/alien_angel_banner.gif`.

No real data is served. Everything is an illusion.

## 🖼️ Examples

- `http://vapornet.net/space_goth_forest/conspiracy_zine04.html`
  A goth teen zine claiming glam rock is a lunar psyop. Includes howling wolf gifs and MIDI of *Baba O'Riley*.

- `http://angelflesh.club/blackmirror_journal/fears4.htm`
  A semi-sentient dream journal about haunted televisions.

- `http://geocities.void/sailor_moon_fan_club14013`
  A sparkly pink shrine to Usagi with pop-up welcome messages and a guestbook that screams when clicked.

## 🧰 Getting Started (Local Dev)

You'll need:

- [Go 1.21+](https://go.dev/)
- An OpenAI API key
- 💻 A haunted CRT (optional but encouraged)

### 1. Clone the repo

```bash
git clone https://github.com/your-user/vapornet.git
cd vapornet
```

### 2. Set up your `.env`

Create a `.env` file in the root of the project with your API key:

```env
OPENAI_API_KEY=sk-...
```

> 🔐 **Pro tip:** Don’t commit this file! Add `.env` to your `.gitignore`.

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the server

```bash
go run main.go
```

This starts your VaporNet dream machine at `http://localhost:8080`.

### 5. Start browsing the net-unreal

Set your browser or HTTP client to use `localhost:8080` as an HTTP proxy (or just `curl` it):

```bash
curl http://ghostforums.biz/alien_truth/files/zeta3.html
```

## 📦 Environment Variables

| Variable         | Description                     |
|------------------|---------------------------------|
| `OPENAI_API_KEY` | Your OpenAI API key (required)  |

## 🔄 Updating Prompt Logic

All prompt templates live in `template/prompt/*.md`.

For example, `concept.md` is the heart of the system—it takes a request and dreams up what kind of page it should become in the world of VaporNet.

> Want to make your hallucinations weirder? Start there.

## 📡 Philosophy

> _"What if we browsed the subconscious instead of the cloud?"_

VaporNet isn’t just a project—it’s a vibe, a resurrection, a love letter to the internet that could’ve been.
It’s emotionally intelligent, dangerously nostalgic, and slightly possessed.

---

**You’re not just browsing. You’re dreaming.**
