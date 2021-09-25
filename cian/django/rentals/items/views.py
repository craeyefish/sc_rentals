from django.http import HttpResponse

def external_home(request):
    return HttpResponse("This is the external home view.")

def home(request):
    return HttpResponse("This is the home view.")

def login(reuquest):
    return HttpResponse("This is the login view.")

def signup(reuquest):
    return HttpResponse("This is the signup view.")
    