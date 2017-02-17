Name: chsht
Version: 0.1.4
Release: 1
Summary: Terminal Cheat Sheet
Requires: less

Group: Applications
License: GPLv2
URL: http://mvrpl.com.br/
Source0: %{name}-%{version}.tar.gz
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root

%description
Terminal Cheat Sheet

%prep
%setup -q

%install
mkdir -p $RPM_BUILD_ROOT/usr/local/bin
install chsht $RPM_BUILD_ROOT/usr/local/bin/chsht

%files
%defattr(-,root,root)
/usr/local/bin/chsht

%clean
rm -rf $RPM_BUILD_ROOT

%changelog
* Tue Oct 11 2016 Marcos Lima <mvrpl@icloud.com> - 0.1.4
- Initial build
