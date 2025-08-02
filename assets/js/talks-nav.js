// Smooth scrolling for year navigation (works for both talks and posts)
document.addEventListener('DOMContentLoaded', function() {
  const yearLinks = document.querySelectorAll('.year-link');
  
  yearLinks.forEach(link => {
    link.addEventListener('click', function(e) {
      e.preventDefault();
      
      const targetId = this.getAttribute('href');
      const targetSection = document.querySelector(targetId);
      
      if (targetSection) {
        // Remove active class from all links
        yearLinks.forEach(l => l.classList.remove('active'));
        
        // Add active class to clicked link
        this.classList.add('active');
        
        // Smooth scroll to section
        targetSection.scrollIntoView({
          behavior: 'smooth',
          block: 'start'
        });
      }
    });
  });
  
  // Highlight active year based on scroll position
  function updateActiveYear() {
    const yearSections = document.querySelectorAll('.year-section');
    const scrollPos = window.scrollY + 100;
    
    yearSections.forEach(section => {
      const sectionTop = section.offsetTop;
      const sectionBottom = sectionTop + section.offsetHeight;
      const yearId = section.getAttribute('id');
      const correspondingLink = document.querySelector(`a[href="#${yearId}"]`);
      
      if (scrollPos >= sectionTop && scrollPos < sectionBottom) {
        yearLinks.forEach(l => l.classList.remove('active'));
        if (correspondingLink) {
          correspondingLink.classList.add('active');
        }
      }
    });
  }
  
  // Update active year on scroll (only if we have year sections)
  if (document.querySelectorAll('.year-section').length > 0) {
    window.addEventListener('scroll', updateActiveYear);
    updateActiveYear();
  }
});
