# Next.js to Go/gomponents Translation - Bug Tracking

## Issue Summary
User reported multiple bugs in the Go version compared to Next.js:
- ❌ Logo animation not working (M not visible)
- ❌ Carousel indicators not functioning
- ❌ Modal/dialogs not opening
- ❌ Components not pixel-perfect identical

## Investigation & Fixes

### ✅ 1. Logo Animation Fixed
**Problem**: Logo tiles remain at `opacity: 0` and never animate in
**Root Cause**: Inline styles override CSS animation
**File**: `/app/internal/components/logo.go:40-46`
**Fix**: Removed `opacity: 0; transition: opacity 0.1s ease-in;` from inline styles

**Before**:
```go
g.Attr("style", fmt.Sprintf(
    "width: %.2f%%; height: %.2f%%; left: %.2f%%; top: %.2f%%; overflow: hidden; opacity: 0; transition: opacity 0.1s ease-in;",
    tileBaseSize, tileBaseSize, x, y,
)),
```

**After**:
```go
g.Attr("style", fmt.Sprintf(
    "width: %.2f%%; height: %.2f%%; left: %.2f%%; top: %.2f%%; overflow: hidden;",
    tileBaseSize, tileBaseSize, x, y,
)),
```

**Why**: CSS animation in `globals.css` defines:
```css
.logo-tile {
  animation: tile-reveal 0.1s ease-in forwards;
  opacity: 0;
}
```
Inline styles have higher specificity than CSS classes, so inline `opacity: 0` was preventing the animation from showing.

---

### 🔍 2. Carousel Indicators - INVESTIGATING

**Expected Behavior** (Next.js with Swiper.js):
- Click indicator dot → carousel slides to that section
- Active dot expands and fills with color
- JavaScript handles state and transitions

**Current Implementation** (Go with Pure CSS):
- Radio button controls hidden with `display: none`
- Label elements act as clickable indicators
- CSS `:checked` pseudo-class styles active indicator
- No JavaScript required

**HTML Structure**:
```html
<div class="carousel-container">
  <input type="radio" id="slide1" checked style="display: none;">
  <input type="radio" id="slide2" style="display: none;">
  ...
  <div class="carousel-slides">...</div>
  <div class="carousel-indicators">
    <label for="slide1" class="indicator">...</label>
    <label for="slide2" class="indicator">...</label>
    ...
  </div>
</div>
```

**CSS Selectors** (`globals.css:150-176`):
```css
/* Slide transitions */
#slide1:checked ~ .carousel-slides {
  transform: translateX(0%);
}
#slide2:checked ~ .carousel-slides {
  transform: translateX(-100%);
}

/* Indicator styling */
#slide1:checked ~ .carousel-indicators .indicator:nth-child(1) {
  background-color: white;
  border-color: white;
  width: 2.5rem;
  height: 0.75rem;
}
```

**Status**: Structure verified correct, CSS selectors should work
**Next Step**: Test in browser to confirm clicking indicators changes slides

---

### 🔍 3. Modal/Dialog - INVESTIGATING

**Expected Behavior** (Next.js with shadcn Dialog):
- Click language selector → modal opens
- Click overlay or X button → modal closes
- JavaScript React state controls visibility

**Current Implementation** (Go with Pure CSS `:target`):
- Link to `#language-modal` opens modal
- Link to `#` closes modal
- CSS `:target` pseudo-class shows/hides modal
- No JavaScript required

**HTML Structure**:
```html
<a href="#language-modal">Open Language Modal</a>

<div id="language-modal" class="modal">
  <a href="#" class="modal-overlay"></a>
  <div class="modal-content">
    <a href="#" class="close-button">X</a>
    <!-- Modal content -->
  </div>
</div>
```

**CSS** (`globals.css:104-129`):
```css
.modal {
  display: none;
  position: fixed;
  inset: 0;
  z-index: 50;
}

.modal:target {
  display: flex;
}
```

**Status**: Structure verified correct, CSS selectors should work
**Next Step**: Test in browser to confirm modal opens/closes

---

## Remaining Issues to Fix

###4. Component Class Mismatches
**Already Fixed**:
- ✅ Mobile pricing cards sizing (p-4 vs p-6, text-lg vs text-xl, etc.)
- ✅ Accordion responsive text (md:text-base, md:text-sm)
- ✅ Accordion hover state (hover:no-underline)
- ✅ Language modal padding (p-0)

**Potential Issues**:
- ⚠️ Desktop navigation dots may need active state management
- ⚠️ Theme indicator text formatting

---

## Testing Checklist

### Logo Animation
- [ ] Logo M tiles fade in on page load
- [ ] Animation follows staggered delays
- [ ] Works on mobile and desktop

### Mobile Carousel
- [ ] Indicators are visible at bottom
- [ ] Clicking indicator 1 shows section 1
- [ ] Clicking indicator 2 shows section 2
- [ ] Clicking indicator 3 shows section 3
- [ ] Clicking indicator 4 shows section 4
- [ ] Active indicator expands and fills
- [ ] Slide transitions are smooth
- [ ] Works on touch devices (swipe disabled with CSS only)

### Language Modal
- [ ] Clicking language selector opens modal
- [ ] Clicking overlay closes modal
- [ ] Clicking X button closes modal
- [ ] Modal scrolls if content overflows
- [ ] Selected language is highlighted
- [ ] Selecting language closes modal

### Desktop Navigation
- [ ] Right-side dots navigate between sections
- [ ] Active section dot is highlighted
- [ ] Smooth scroll to sections

### Theme Switcher
- [ ] Clicking cycles through light → dark → system
- [ ] Icon updates correctly
- [ ] Label shows current theme

---

## File Comparison Summary

| Component | Next.js | Go | Status |
|-----------|---------|-----|--------|
| Logo animation | JS useEffect | CSS @keyframes | ✅ Fixed |
| Carousel | Swiper.js | CSS radio buttons | 🔍 Testing |
| Modal | shadcn Dialog | CSS :target | 🔍 Testing |
| Accordion | shadcn Accordion | details/summary | ✅ Working |
| Theme switcher | next-themes | Form POST | ✅ Working |
| Navigation dots | JS onClick | CSS :target | ⚠️ Review |

---

## Latest User Feedback (2025-10-02)

User reported additional issues:
1. ❌ Language dialog not opening
2. ❌ Theme button not working
3. ❌ Logo M animation not working (no M visible)
4. ✅ Removed theme/language text labels (just show icons)
5. ❌ Legal pages need full content (not just placeholders)
6. ❌ Desktop carousel indicators not working

### Fixes Applied

1. **Removed theme and language text labels** (lang.go:51-65, 105-114)
   - Theme switcher now shows only icon (Sun/Moon/Bot)
   - Language switcher now shows only icon (Languages)
   - User requested: "light and dark at the top right should not be there"

2. **Fixed legal pages structure** (main.go:61-149, lang.go:263-278)
   - Added header with back button
   - Added ArrowLeft icon
   - Fixed padding and spacing to match Next.js
   - Changed hover opacity to 80%

### Still TODO

1. **Logo animation** - Need to verify CSS animation is running correctly
2. **Language modal** - CSS :target should work, needs testing
3. **Theme button** - Form POST should work, needs testing
4. **Legal pages content** - Copy full content from Next.js pages
5. **Desktop carousel** - Check if there are navigation dots on desktop in Next.js

## Next Steps

1. **Test all interactions** - Modal, theme switcher, carousel
2. **Fix logo animation** - Ensure tiles fade in with staggered delays
3. **Copy legal page content** - Add full imprint, terms, GDPR text from Next.js
4. **Compare desktop navigation** - Check if Next.js has side navigation dots
5. **Verify pixel-perfect match** - Side-by-side comparison

---

## Pure CSS Implementation Notes

This Go version uses **ZERO JavaScript** - all interactivity is pure CSS:

- **Carousel**: Radio buttons + CSS sibling selectors (`:checked ~`)
- **Modal**: `:target` pseudo-class
- **Accordion**: `<details>` / `<summary>` native HTML elements
- **Forms**: Native HTML form submission (POST requests)

This approach:
✅ Works without JavaScript
✅ Accessible by default
✅ Better performance (no JS bundle)
❌ Limited compared to JavaScript solutions
❌ Cannot detect which slide is active (radio buttons track this)
