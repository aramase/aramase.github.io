// Cool animations and interactive effects
document.addEventListener('DOMContentLoaded', function() {
    
    // Enhanced Intersection Observer for scroll animations
    const observerOptions = {
        threshold: 0.15,
        rootMargin: '0px 0px -100px 0px'
    };

    const observer = new IntersectionObserver((entries) => {
        entries.forEach((entry, index) => {
            if (entry.isIntersecting) {
                // Stagger animation delays for multiple elements
                setTimeout(() => {
                    entry.target.classList.add('animate-in');
                }, index * 100);
            }
        });
    }, observerOptions);

    // Observe elements for scroll animations with better selectors
    document.querySelectorAll('.talk-item, .blog-card, .post-card, .about-intro, .social-icons').forEach(el => {
        observer.observe(el);
    });

    // Enhanced hover effects for cards - less intrusive
    document.querySelectorAll('.talk-item, .blog-card, .post-card').forEach(card => {
        card.addEventListener('mouseenter', function() {
            // Use CSS class instead of direct style manipulation
            this.classList.add('card-hover');
        });
        
        card.addEventListener('mouseleave', function() {
            this.classList.remove('card-hover');
        });
    });

    // Interactive talk previews with YouTube integration
    function setupTalkPreviews() {
        document.querySelectorAll('.talk-image a').forEach(link => {
            const videoUrl = link.href;
            const isYouTube = videoUrl.includes('youtube.com') || videoUrl.includes('youtu.be');
            
            if (isYouTube) {
                const videoId = extractYouTubeId(videoUrl);
                if (videoId) {
                    addInteractiveThumbnail(link, videoId);
                }
            }
        });
    }

    function extractYouTubeId(url) {
        const regExp = /^.*(youtu.be\/|v\/|u\/\w\/|embed\/|watch\?v=|&v=)([^#&?]*).*/;
        const match = url.match(regExp);
        return (match && match[2].length === 11) ? match[2] : null;
    }

    function addInteractiveThumbnail(link, videoId) {
        const thumbnail = link.querySelector('img');
        if (thumbnail) {
            // Replace with high-quality YouTube thumbnail
            thumbnail.src = `https://img.youtube.com/vi/${videoId}/maxresdefault.jpg`;
            thumbnail.onerror = function() {
                this.src = `https://img.youtube.com/vi/${videoId}/hqdefault.jpg`;
            };
            
            // Add enhanced play button
            const playButton = link.querySelector('.play-button') || document.createElement('div');
            playButton.className = 'play-button enhanced';
            playButton.innerHTML = `
                <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M8 5v14l11-7z"/>
                </svg>
            `;
            
            if (!link.querySelector('.play-button')) {
                link.appendChild(playButton);
            }
            
            // Add hover preview effect
            link.addEventListener('mouseenter', function() {
                const preview = createVideoPreview(videoId);
                this.appendChild(preview);
                setTimeout(() => preview.classList.add('show'), 100);
            });
            
            link.addEventListener('mouseleave', function() {
                const preview = this.querySelector('.video-preview');
                if (preview) {
                    preview.classList.remove('show');
                    setTimeout(() => preview.remove(), 300);
                }
            });
        }
    }

    function createVideoPreview(videoId) {
        const preview = document.createElement('div');
        preview.className = 'video-preview';
        preview.innerHTML = `
            <div class="preview-content">
                <div class="preview-title">Click to watch</div>
                <div class="preview-duration" data-video-id="${videoId}">Loading...</div>
            </div>
        `;
        return preview;
    }

    // Back to Top Button Setup
    function setupBackToTop() {
        const backToTopButton = document.getElementById('back-to-top');
        if (!backToTopButton) return;

        let isVisible = false;
        let ticking = false;

        function updateButtonVisibility() {
            const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
            const shouldShow = scrollTop > 300; // Show after scrolling 300px

            if (shouldShow && !isVisible) {
                backToTopButton.classList.add('visible');
                isVisible = true;
            } else if (!shouldShow && isVisible) {
                backToTopButton.classList.remove('visible');
                isVisible = false;
            }
            ticking = false;
        }

        function onScroll() {
            if (!ticking) {
                requestAnimationFrame(updateButtonVisibility);
                ticking = true;
            }
        }

        function smoothScrollToTop() {
            const startTime = performance.now();
            const startPosition = window.pageYOffset || document.documentElement.scrollTop;
            const duration = 800; // 800ms for smooth scroll

            function scrollAnimation(currentTime) {
                const elapsed = currentTime - startTime;
                const progress = Math.min(elapsed / duration, 1);
                
                // Easing function for smooth deceleration
                const easeOutQuart = 1 - Math.pow(1 - progress, 4);
                
                window.scrollTo(0, startPosition * (1 - easeOutQuart));

                if (progress < 1) {
                    requestAnimationFrame(scrollAnimation);
                }
            }

            requestAnimationFrame(scrollAnimation);
        }

        // Event listeners
        window.addEventListener('scroll', onScroll, { passive: true });
        backToTopButton.addEventListener('click', smoothScrollToTop);

        // Initial check
        updateButtonVisibility();
    }

    // Initialize talk previews
    setupTalkPreviews();

    // Back to Top Button functionality
    setupBackToTop();

    // Typing effect for homepage title
    function typeWriter(element, text, speed = 100) {
        let i = 0;
        element.innerHTML = '';
        
        function type() {
            if (i < text.length) {
                element.innerHTML += text.charAt(i);
                i++;
                setTimeout(type, speed);
            }
        }
        type();
    }

    // Apply typing effect to main title if on homepage
    const mainTitle = document.querySelector('.about-intro h1');
    if (mainTitle && window.location.pathname === '/') {
        const originalText = mainTitle.textContent;
        typeWriter(mainTitle, originalText, 150);
    }

    // Smooth scrolling for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                target.scrollIntoView({
                    behavior: 'smooth',
                    block: 'start'
                });
            }
        });
    });

});
