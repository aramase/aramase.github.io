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

    // Add hover effects to cards
    document.querySelectorAll('.talk, .blog-post').forEach(card => {
        card.addEventListener('mouseenter', function() {
            this.style.transform = 'translateY(-5px)';
            this.style.boxShadow = '0 10px 30px rgba(59, 130, 246, 0.2)';
        });

        card.addEventListener('mouseleave', function() {
            this.style.transform = 'translateY(0)';
            this.style.boxShadow = 'none';
        });
    });

    // Parallax effect for profile image
    const profileImage = document.querySelector('.about-image');
    if (profileImage) {
        window.addEventListener('scroll', () => {
            const scrolled = window.pageYOffset;
            const rate = scrolled * -0.2;
            profileImage.style.transform = `translateY(${rate}px)`;
        });
    }

    // Add loading animation
    document.body.classList.add('loaded');
});

// CSS class for enhanced scroll animations and interactive elements
const style = document.createElement('style');
style.textContent = `
    /* Enhanced scroll animations */
    .animate-in {
        animation: slideInUp 0.8s cubic-bezier(0.25, 0.46, 0.45, 0.94);
    }

    @keyframes slideInUp {
        from {
            opacity: 0;
            transform: translateY(40px) scale(0.95);
        }
        to {
            opacity: 1;
            transform: translateY(0) scale(1);
        }
    }

    /* Enhanced card transitions */
    .talk-item, .blog-card, .post-card {
        transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
        will-change: transform, box-shadow;
    }

    /* Enhanced play button */
    .play-button.enhanced {
        background: rgba(0, 0, 0, 0.8);
        backdrop-filter: blur(10px);
        border: 2px solid rgba(255, 255, 255, 0.3);
        width: 70px;
        height: 70px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%) scale(0.9);
        transition: all 0.3s ease;
        z-index: 2;
    }

    .talk-image:hover .play-button.enhanced {
        transform: translate(-50%, -50%) scale(1);
        background: rgba(96, 165, 250, 0.9);
        border-color: rgba(255, 255, 255, 0.5);
    }

    /* Video preview tooltip */
    .video-preview {
        position: absolute;
        bottom: 100%;
        left: 50%;
        transform: translateX(-50%);
        background: rgba(0, 0, 0, 0.9);
        color: white;
        padding: 12px 16px;
        border-radius: 8px;
        margin-bottom: 10px;
        opacity: 0;
        transform: translateX(-50%) translateY(10px);
        transition: all 0.3s ease;
        pointer-events: none;
        z-index: 10;
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.1);
    }

    .video-preview.show {
        opacity: 1;
        transform: translateX(-50%) translateY(0);
    }

    .video-preview::after {
        content: '';
        position: absolute;
        top: 100%;
        left: 50%;
        transform: translateX(-50%);
        border: 6px solid transparent;
        border-top-color: rgba(0, 0, 0, 0.9);
    }

    .preview-title {
        font-size: 14px;
        font-weight: 600;
        margin-bottom: 4px;
    }

    .preview-duration {
        font-size: 12px;
        opacity: 0.8;
    }

    /* Smooth loading animation */
    body {
        opacity: 0;
        transition: opacity 0.5s ease;
    }

    body.loaded {
        opacity: 1;
    }

    /* Subtle glow effect for interactive elements */
    .talk-item:hover .talk-thumbnail,
    .blog-card:hover,
    .post-card:hover {
        box-shadow: 0 0 30px rgba(96, 165, 250, 0.2);
    }

    /* Enhanced social icons animation */
    .social-icons a {
        transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
    }

    .social-icons a:hover {
        transform: translateY(-3px) scale(1.1);
        color: var(--accent);
    }

    /* Staggered animation for multiple elements */
    .talk-item:nth-child(1) { animation-delay: 0.1s; }
    .talk-item:nth-child(2) { animation-delay: 0.2s; }
    .talk-item:nth-child(3) { animation-delay: 0.3s; }
    .talk-item:nth-child(4) { animation-delay: 0.4s; }
    
    .blog-card:nth-child(1) { animation-delay: 0.1s; }
    .blog-card:nth-child(2) { animation-delay: 0.2s; }
    .blog-card:nth-child(3) { animation-delay: 0.3s; }
`;
document.head.appendChild(style);
