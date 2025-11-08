const CACHE_NAME = 'lomboktojson-v1';
const RUNTIME_CACHE = 'lomboktojson-runtime-v1';

// Assets to cache on install
const STATIC_ASSETS = [
  './',
  './index.html',
  './manifest.json',
  './favicon.ico',
  './assets/lombok2json.wasm',
  './assets/wasm_exec.js',
  './assets/app.js',
  './assets/icon-48.png',
  './assets/icon-72.png',
  './assets/icon-96.png',
  './assets/icon-144.png',
  './assets/icon-192.png',
  './assets/icon-512.png',
  './assets/high-level-idea-clear.png',
  './assets/lomboktojson-banner.png',
  './assets/website-screenshot.png'
];

// External CDN resources to cache on first access
const CDN_RESOURCES = [
  'https://cdnjs.cloudflare.com/ajax/libs/tailwindcss/2.2.19/tailwind.min.css',
  'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.34.0/min/vs/loader.min.js',
  'https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap'
];

// Install event - cache static assets
self.addEventListener('install', event => {
  console.log('[Service Worker] Installing...');
  event.waitUntil(
    caches.open(CACHE_NAME)
      .then(cache => {
        console.log('[Service Worker] Caching static assets');
        return cache.addAll(STATIC_ASSETS);
      })
      .then(() => self.skipWaiting())
  );
});

// Activate event - clean up old caches
self.addEventListener('activate', event => {
  console.log('[Service Worker] Activating...');
  event.waitUntil(
    caches.keys().then(cacheNames => {
      return Promise.all(
        cacheNames
          .filter(name => name !== CACHE_NAME && name !== RUNTIME_CACHE)
          .map(name => {
            console.log('[Service Worker] Deleting old cache:', name);
            return caches.delete(name);
          })
      );
    }).then(() => self.clients.claim())
  );
});

// Fetch event - serve from cache, fallback to network
self.addEventListener('fetch', event => {
  const { request } = event;
  const url = new URL(request.url);

  // Skip caching for Google Analytics
  if (url.hostname === 'www.googletagmanager.com' || url.hostname === 'www.google-analytics.com') {
    return;
  }

  // For WASM files, use cache-first strategy
  if (request.url.endsWith('.wasm')) {
    event.respondWith(
      caches.match(request)
        .then(response => {
          if (response) {
            return response;
          }
          return fetch(request).then(response => {
            return caches.open(CACHE_NAME).then(cache => {
              cache.put(request, response.clone());
              return response;
            });
          });
        })
    );
    return;
  }

  // For CDN resources and fonts, use cache-first strategy
  if (url.hostname.includes('cdnjs.cloudflare.com') ||
      url.hostname.includes('fonts.googleapis.com') ||
      url.hostname.includes('fonts.gstatic.com')) {
    event.respondWith(
      caches.match(request)
        .then(response => {
          if (response) {
            return response;
          }
          return fetch(request).then(response => {
            // Only cache successful responses
            if (response && response.status === 200) {
              return caches.open(RUNTIME_CACHE).then(cache => {
                cache.put(request, response.clone());
                return response;
              });
            }
            return response;
          });
        })
        .catch(() => {
          // Return offline fallback if available
          console.log('[Service Worker] Fetch failed for:', request.url);
        })
    );
    return;
  }

  // For Monaco Editor resources
  if (url.pathname.includes('/monaco-editor/')) {
    event.respondWith(
      caches.match(request)
        .then(response => {
          if (response) {
            return response;
          }
          return fetch(request).then(response => {
            if (response && response.status === 200) {
              return caches.open(RUNTIME_CACHE).then(cache => {
                cache.put(request, response.clone());
                return response;
              });
            }
            return response;
          });
        })
    );
    return;
  }

  // For all other requests, try cache first, then network
  event.respondWith(
    caches.match(request)
      .then(response => {
        if (response) {
          return response;
        }
        return fetch(request).then(response => {
          // Check if we received a valid response
          if (!response || response.status !== 200 || response.type !== 'basic') {
            return response;
          }

          // Clone the response
          const responseToCache = response.clone();

          caches.open(RUNTIME_CACHE)
            .then(cache => {
              cache.put(request, responseToCache);
            });

          return response;
        });
      })
      .catch(() => {
        // If both cache and network fail, return offline page
        if (request.destination === 'document') {
          return caches.match('./index.html');
        }
      })
  );
});

// Handle messages from clients
self.addEventListener('message', event => {
  if (event.data && event.data.type === 'SKIP_WAITING') {
    self.skipWaiting();
  }
});
